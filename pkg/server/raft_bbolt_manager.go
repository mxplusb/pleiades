/*
 * Copyright (c) 2022 Sienna Lloyd
 *
 * Licensed under the PolyForm Strict License 1.0.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License here:
 *  https://github.com/mxplusb/pleiades/blob/mainline/LICENSE
 */

package server

import (
	"context"
	"time"

	"github.com/mxplusb/pleiades/api/v1/database"
	"github.com/mxplusb/pleiades/pkg/fsm/kv"
	"github.com/mxplusb/pleiades/pkg/routing"
	"github.com/cockroachdb/errors"
	"github.com/lni/dragonboat/v3"
	"github.com/lni/dragonboat/v3/client"
	"github.com/rs/zerolog"
)

var (
	_ IStore = (*bboltStoreManager)(nil)
)

func newBboltStoreManager(tm *raftTransactionManager, nh *dragonboat.NodeHost, logger zerolog.Logger) *bboltStoreManager {
	l := logger.With().Str("component", "store-manager").Logger()
	return &bboltStoreManager{l, tm, nh, &routing.Shard{}}
}

type bboltStoreManager struct {
	logger zerolog.Logger
	tm     *raftTransactionManager
	nh     *dragonboat.NodeHost
	shardRouter *routing.Shard
}

func (s *bboltStoreManager) CreateAccount(request *database.CreateAccountRequest) (*database.CreateAccountReply, error) {

	account := request.GetAccountId()
	if account == 0 {
		s.logger.Trace().Msg("empty account value")
		return &database.CreateAccountReply{}, kv.ErrInvalidAccount
	}

	owner := request.GetOwner()
	if owner == "" {
		s.logger.Trace().Msg("empty owner value")
		return &database.CreateAccountReply{}, kv.ErrInvalidOwner
	}

	req := &database.KVStoreWrapper{
		Account: request.GetAccountId(),
		Typ:     database.KVStoreWrapper_CREATE_ACCOUNT_REQUEST,
		Payload: &database.KVStoreWrapper_CreateAccountRequest{
			CreateAccountRequest: request,
		},
	}

	cmd, err := req.MarshalVT()
	if err != nil {
		s.logger.Error().Err(err).Msg("can't marshal request")
		return &database.CreateAccountReply{}, errors.Wrap(err, "can't marshal request")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()

	var cs *client.Session
	if request.Transaction != nil {
		cs = s.tm.sessionCache[request.GetTransaction().GetClientId()]
	} else {
		shard := s.shardRouter.AccountToShard(account)
		cs = s.nh.GetNoOPSession(shard)
	}

	result, err := s.nh.SyncPropose(ctx, cs, cmd)
	if err != nil {
		s.logger.Error().Err(err).Msg("can't apply message")
		return &database.CreateAccountReply{}, errors.Wrap(err, "can't apply message")
	}

	if cs.SeriesID != client.NoOPSeriesID {
		cs.ProposalCompleted()
	}

	resp := &database.KVStoreWrapper{}
	err = resp.UnmarshalVT(result.Data)
	if err != nil {
		s.logger.Error().Err(err).Msg("can't unmarshal response")
		return &database.CreateAccountReply{}, errors.Wrap(err, "can't unmarshal response")
	}

	response := &database.CreateAccountReply{}
	if resp.GetCreateAccountReply() == nil {
		response.AccountDescriptor = &database.AccountDescriptor{}
	} else {
		response.AccountDescriptor = resp.GetCreateAccountReply().GetAccountDescriptor()
	}

	if request.Transaction == nil {
		response.Transaction = &database.Transaction{}
	} else {
		response.Transaction = csToTransaction(*cs)
	}

	return response, nil
}

func (s *bboltStoreManager) DeleteAccount(request *database.DeleteBucketRequest) (*database.DeleteAccountReply, error) {
	//TODO implement me
	panic("implement me")
}

func (s *bboltStoreManager) CreateBucket(request *database.CreateBucketRequest) (*database.CreateBucketReply, error) {
	//TODO implement me
	panic("implement me")
}

func (s *bboltStoreManager) DeleteBucket(request *database.DeleteBucketRequest) (*database.DeleteBucketReply, error) {
	//TODO implement me
	panic("implement me")
}

func (s *bboltStoreManager) GetKey(request *database.GetKeyRequest) (*database.GetKeyReply, error) {
	//TODO implement me
	panic("implement me")
}

func (s *bboltStoreManager) PutKey(request *database.PutKeyRequest) (*database.PutKeyReply, error) {
	//TODO implement me
	panic("implement me")
}

func (s *bboltStoreManager) DeleteKey(request *database.DeleteKeyRequest) (*database.DeleteKeyReply, error) {
	//TODO implement me
	panic("implement me")
}
