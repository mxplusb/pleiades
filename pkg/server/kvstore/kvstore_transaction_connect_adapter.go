/*
 * Copyright (c) 2022-2023 Sienna Lloyd
 *
 * Licensed under the PolyForm Internal Use License 1.0.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License here:
 *  https://github.com/mxplusb/pleiades/blob/mainline/LICENSE
 */

package kvstore

import (
	"context"
	"net/http"

	kvstorev1 "github.com/mxplusb/pleiades/pkg/api/kvstore/v1"
	"github.com/mxplusb/pleiades/pkg/api/kvstore/v1/kvstorev1connect"
	"github.com/mxplusb/pleiades/pkg/server/runtime"
	"github.com/mxplusb/pleiades/pkg/server/transactions"
	"github.com/bufbuild/connect-go"
	"github.com/cockroachdb/errors"
	dclient "github.com/lni/dragonboat/v3/client"
	"github.com/rs/zerolog"
	"go.uber.org/fx"
)

var (
	_ kvstorev1connect.TransactionsServiceHandler = (*KvStoreTransactionConnectAdapter)(nil)
	_ runtime.ServiceHandler                      = (*KvStoreTransactionConnectAdapter)(nil)
)

type KvStoreTransactionConnectAdapterBuilderParams struct {
	fx.In
	TransactionManager runtime.ITransactionManager
	Logger             zerolog.Logger
}

type KvStoreTransactionConnectAdapterBuilderResults struct {
	fx.In

	ConnectAdapter *KvStoreTransactionConnectAdapter
}

type KvStoreTransactionConnectAdapter struct {
	http.Handler
	logger             zerolog.Logger
	transactionManager runtime.ITransactionManager
	path               string
}

func NewKvstoreTransactionConnectAdapter(transactionManager runtime.ITransactionManager, logger zerolog.Logger) *KvStoreTransactionConnectAdapter {
	adapter := &KvStoreTransactionConnectAdapter{logger: logger, transactionManager: transactionManager}
	adapter.path, adapter.Handler = kvstorev1connect.NewTransactionsServiceHandler(adapter)
	return adapter
}

func (k *KvStoreTransactionConnectAdapter) Path() string {
	return k.path
}

func (k *KvStoreTransactionConnectAdapter) NewTransaction(ctx context.Context, c *connect.Request[kvstorev1.NewTransactionRequest]) (*connect.Response[kvstorev1.NewTransactionResponse], error) {
	if c.Msg.GetShardId() == 0 {
		return connect.NewResponse(&kvstorev1.NewTransactionResponse{}), errors.New("shard id must not be 0")
	}

	transaction, err := k.transactionManager.GetTransaction(ctx, c.Msg.GetShardId())
	if err != nil {
		k.logger.Error().Err(err).Uint64("shard", c.Msg.GetShardId()).Msg("cannot create transaction")
	}
	return connect.NewResponse(&kvstorev1.NewTransactionResponse{Transaction: transaction}), nil
}

func (k *KvStoreTransactionConnectAdapter) CloseTransaction(ctx context.Context, c *connect.Request[kvstorev1.CloseTransactionRequest]) (*connect.Response[kvstorev1.CloseTransactionResponse], error) {
	transaction := c.Msg.GetTransaction()
	if err := k.checkTransaction(transaction); err != nil {
		return connect.NewResponse(&kvstorev1.CloseTransactionResponse{}), err
	}

	err := k.transactionManager.CloseTransaction(ctx, transaction)
	if err != nil {
		k.logger.Error().Err(err).Msg("can't close transaction")
	}
	return connect.NewResponse(&kvstorev1.CloseTransactionResponse{}), err
}

func (k *KvStoreTransactionConnectAdapter) Commit(ctx context.Context, c *connect.Request[kvstorev1.CommitRequest]) (*connect.Response[kvstorev1.CommitResponse], error) {
	transaction := c.Msg.GetTransaction()
	if err := k.checkTransaction(transaction); err != nil {
		return connect.NewResponse(&kvstorev1.CommitResponse{}), err
	}

	t := k.transactionManager.Commit(ctx, transaction)

	return connect.NewResponse(&kvstorev1.CommitResponse{Transaction: t}), nil
}

// todo (sienna): replace this with dclient.Session.ValidForProposal later.
func (k *KvStoreTransactionConnectAdapter) checkTransaction(t *kvstorev1.Transaction) error {
	// I don't think this can happen because it's a pointer, but better to be safe than sorry
	if t == nil {
		k.logger.Error().Err(transactions.ErrNilTransaction).Msg("attempted close of an empty transaction")
		return transactions.ErrNilTransaction
	}

	// check for noop or unset
	if t.GetTransactionId() == dclient.NoOPSeriesID {
		return transactions.ErrUnupportedTransaction
	}

	// check for unregister
	if t.GetTransactionId() == dclient.SeriesIDForUnregister {
		return transactions.ErrUnupportedTransaction
	}

	// check for pending registration
	if t.TransactionId == dclient.SeriesIDForRegister {
		return transactions.ErrUnupportedTransaction
	}

	return nil
}
