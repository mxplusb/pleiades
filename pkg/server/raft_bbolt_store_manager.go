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

	aerrs "github.com/mxplusb/api/errors/v1"
	kvstorev1 "github.com/mxplusb/api/kvstore/v1"
	"github.com/mxplusb/pleiades/pkg/fsm/kv"
	"github.com/mxplusb/pleiades/pkg/routing"
	"github.com/mxplusb/pleiades/pkg/utils"
	"github.com/cockroachdb/errors"
	"github.com/lni/dragonboat/v3"
	"github.com/lni/dragonboat/v3/client"
	"github.com/rs/zerolog"
)

var (
	_ IKVStore = (*bboltStoreManager)(nil)
)

func newBboltStoreManager(tm *raftTransactionManager, nh *dragonboat.NodeHost, logger zerolog.Logger) *bboltStoreManager {
	l := logger.With().Str("component", "store-manager").Logger()
	return &bboltStoreManager{l, tm, nh, &routing.Shard{}, 1000 * time.Millisecond}
}

type bboltStoreManager struct {
	logger         zerolog.Logger
	tm             *raftTransactionManager
	nh             *dragonboat.NodeHost
	shardRouter    *routing.Shard
	defaultTimeout time.Duration
}

func (s *bboltStoreManager) CreateAccount(request *kvstorev1.CreateAccountRequest) (*kvstorev1.CreateAccountResponse, error) {

	account := request.GetAccountId()
	if account == 0 {
		s.logger.Trace().Msg("empty account value")
		return &kvstorev1.CreateAccountResponse{}, kv.ErrInvalidAccount
	}

	owner := request.GetOwner()
	if owner == "" {
		s.logger.Trace().Msg("empty owner value")
		return &kvstorev1.CreateAccountResponse{}, kv.ErrInvalidOwner
	}

	req := &kvstorev1.KVStoreWrapper{
		Account: request.GetAccountId(),
		Typ:     kvstorev1.KVStoreWrapper_REQUEST_TYPE_CREATE_ACCOUNT_REQUEST,
		Payload: &kvstorev1.KVStoreWrapper_CreateAccountRequest{
			CreateAccountRequest: request,
		},
	}

	cmd, err := req.MarshalVT()
	if err != nil {
		s.logger.Error().Err(err).Msg("can't marshal request")
		return &kvstorev1.CreateAccountResponse{}, errors.Wrap(err, "can't marshal request")
	}

	ctx, cancel := context.WithTimeout(context.Background(), utils.Timeout(s.defaultTimeout))
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
		return &kvstorev1.CreateAccountResponse{}, errors.Wrap(err, "can't apply message")
	}

	if cs.SeriesID != client.NoOPSeriesID {
		cs.ProposalCompleted()
	}

	resp := &kvstorev1.KVStoreWrapper{}
	err = resp.UnmarshalVT(result.Data)
	if err != nil {
		s.logger.Error().Err(err).Msg("can't unmarshal response")
		return &kvstorev1.CreateAccountResponse{}, errors.Wrap(err, "can't unmarshal response")
	}

	if resp.Typ == kvstorev1.KVStoreWrapper_REQUEST_TYPE_RECOVERABLE_ERROR {
		errMsg := &aerrs.Error{}
		errMsg = resp.GetError()

		return &kvstorev1.CreateAccountResponse{}, errors.New(errMsg.Message)
	}

	response := &kvstorev1.CreateAccountResponse{}
	if resp.GetCreateAccountReply() == nil {
		response.AccountDescriptor = &kvstorev1.AccountDescriptor{}
	} else {
		response.AccountDescriptor = resp.GetCreateAccountReply().GetAccountDescriptor()
	}

	if request.Transaction == nil {
		response.Transaction = &kvstorev1.Transaction{}
	} else {
		response.Transaction = csToTransaction(*cs)
	}

	return response, nil
}

func (s *bboltStoreManager) DeleteAccount(request *kvstorev1.DeleteAccountRequest) (*kvstorev1.DeleteAccountResponse, error) {
	account := request.GetAccountId()
	if account == 0 {
		s.logger.Trace().Msg("empty account value")
		return &kvstorev1.DeleteAccountResponse{}, kv.ErrInvalidAccount
	}

	owner := request.GetOwner()
	if owner == "" {
		s.logger.Trace().Msg("empty owner value")
		return &kvstorev1.DeleteAccountResponse{}, kv.ErrInvalidOwner
	}

	req := &kvstorev1.KVStoreWrapper{
		Account: request.GetAccountId(),
		Typ:     kvstorev1.KVStoreWrapper_REQUEST_TYPE_DELETE_ACCOUNT_REQUEST,
		Payload: &kvstorev1.KVStoreWrapper_DeleteAccountRequest{
			DeleteAccountRequest: request,
		},
	}

	cmd, err := req.MarshalVT()
	if err != nil {
		s.logger.Error().Err(err).Msg("can't marshal request")
		return &kvstorev1.DeleteAccountResponse{}, errors.Wrap(err, "can't marshal request")
	}

	ctx, cancel := context.WithTimeout(context.Background(), utils.Timeout(s.defaultTimeout))
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
		return &kvstorev1.DeleteAccountResponse{}, errors.Wrap(err, "can't apply message")
	}

	if cs.SeriesID != client.NoOPSeriesID {
		cs.ProposalCompleted()
	}

	resp := &kvstorev1.KVStoreWrapper{}
	err = resp.UnmarshalVT(result.Data)
	if err != nil {
		s.logger.Error().Err(err).Msg("can't unmarshal response")
		return &kvstorev1.DeleteAccountResponse{}, errors.Wrap(err, "can't unmarshal response")
	}

	if resp.Typ == kvstorev1.KVStoreWrapper_REQUEST_TYPE_RECOVERABLE_ERROR {
		errMsg := &aerrs.Error{}
		errMsg = resp.GetError()

		return &kvstorev1.DeleteAccountResponse{}, errors.New(errMsg.Message)
	}

	response := &kvstorev1.DeleteAccountResponse{
		Ok: resp.GetDeleteAccountReply().GetOk(),
	}

	if request.Transaction == nil || cs.SeriesID != client.NoOPSeriesID {
		response.Transaction = &kvstorev1.Transaction{}
	} else {
		response.Transaction = csToTransaction(*cs)
	}

	return response, nil
}

func (s *bboltStoreManager) CreateBucket(request *kvstorev1.CreateBucketRequest) (*kvstorev1.CreateBucketResponse, error) {
	account := request.GetAccountId()
	if account == 0 {
		s.logger.Trace().Msg("empty account value")
		return &kvstorev1.CreateBucketResponse{}, kv.ErrInvalidAccount
	}

	owner := request.GetOwner()
	if owner == "" {
		s.logger.Trace().Msg("empty owner value")
		return &kvstorev1.CreateBucketResponse{}, kv.ErrInvalidOwner
	}

	bucketName := request.GetName()
	if bucketName == "" {
		s.logger.Trace().Msg("empty bucket name")
		return &kvstorev1.CreateBucketResponse{}, kv.ErrInvalidBucketName
	}

	req := &kvstorev1.KVStoreWrapper{
		Account: request.GetAccountId(),
		Bucket:  bucketName,
		Typ:     kvstorev1.KVStoreWrapper_REQUEST_TYPE_CREATE_BUCKET_REQUEST,
		Payload: &kvstorev1.KVStoreWrapper_CreateBucketRequest{
			CreateBucketRequest: request,
		},
	}

	cmd, err := req.MarshalVT()
	if err != nil {
		s.logger.Error().Err(err).Msg("can't marshal request")
		return &kvstorev1.CreateBucketResponse{}, errors.Wrap(err, "can't marshal request")
	}

	ctx, cancel := context.WithTimeout(context.Background(), utils.Timeout(s.defaultTimeout))
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
		return &kvstorev1.CreateBucketResponse{}, errors.Wrap(err, "can't apply message")
	}

	if cs.SeriesID != client.NoOPSeriesID {
		cs.ProposalCompleted()
	}

	resp := &kvstorev1.KVStoreWrapper{}
	err = resp.UnmarshalVT(result.Data)
	if err != nil {
		s.logger.Error().Err(err).Msg("can't unmarshal response")
		return &kvstorev1.CreateBucketResponse{}, errors.Wrap(err, "can't unmarshal response")
	}

	if resp.Typ == kvstorev1.KVStoreWrapper_REQUEST_TYPE_RECOVERABLE_ERROR {
		errMsg := &aerrs.Error{}
		errMsg = resp.GetError()

		return &kvstorev1.CreateBucketResponse{}, errors.New(errMsg.Message)
	}

	response := &kvstorev1.CreateBucketResponse{}
	if resp.GetCreateBucketReply() == nil {
		response.BucketDescriptor = &kvstorev1.BucketDescriptor{}
	} else {
		response.BucketDescriptor = resp.GetCreateBucketReply().GetBucketDescriptor()
	}

	if request.Transaction == nil {
		response.Transaction = &kvstorev1.Transaction{}
	} else {
		response.Transaction = csToTransaction(*cs)
	}

	return response, nil
}

func (s *bboltStoreManager) DeleteBucket(request *kvstorev1.DeleteBucketRequest) (*kvstorev1.DeleteBucketResponse, error) {
	account := request.GetAccountId()
	if account == 0 {
		s.logger.Trace().Msg("empty account value")
		return &kvstorev1.DeleteBucketResponse{}, kv.ErrInvalidAccount
	}

	name := request.GetName()
	if name == "" {
		s.logger.Trace().Msg("empty name value")
		return &kvstorev1.DeleteBucketResponse{}, kv.ErrInvalidOwner
	}

	req := &kvstorev1.KVStoreWrapper{
		Account: request.GetAccountId(),
		Bucket:  name,
		Typ:     kvstorev1.KVStoreWrapper_REQUEST_TYPE_DELETE_BUCKET_REQUEST,
		Payload: &kvstorev1.KVStoreWrapper_DeleteBucketRequest{
			DeleteBucketRequest: request,
		},
	}

	cmd, err := req.MarshalVT()
	if err != nil {
		s.logger.Error().Err(err).Msg("can't marshal request")
		return &kvstorev1.DeleteBucketResponse{}, errors.Wrap(err, "can't marshal request")
	}

	ctx, cancel := context.WithTimeout(context.Background(), utils.Timeout(s.defaultTimeout))
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
		return &kvstorev1.DeleteBucketResponse{}, errors.Wrap(err, "can't apply message")
	}

	if cs.SeriesID != client.NoOPSeriesID {
		cs.ProposalCompleted()
	}

	resp := &kvstorev1.KVStoreWrapper{}
	err = resp.UnmarshalVT(result.Data)
	if err != nil {
		s.logger.Error().Err(err).Msg("can't unmarshal response")
		return &kvstorev1.DeleteBucketResponse{}, errors.Wrap(err, "can't unmarshal response")
	}

	if resp.Typ == kvstorev1.KVStoreWrapper_REQUEST_TYPE_RECOVERABLE_ERROR {
		errMsg := &aerrs.Error{}
		errMsg = resp.GetError()

		return &kvstorev1.DeleteBucketResponse{}, errors.New(errMsg.Message)
	}

	response := &kvstorev1.DeleteBucketResponse{
		Ok: resp.GetDeleteBucketReply().GetOk(),
	}

	if request.Transaction == nil || cs.SeriesID != client.NoOPSeriesID {
		response.Transaction = &kvstorev1.Transaction{}
	} else {
		response.Transaction = csToTransaction(*cs)
	}

	return response, nil
}

func (s *bboltStoreManager) GetKey(request *kvstorev1.GetKeyRequest) (*kvstorev1.GetKeyResponse, error) {
	account := request.GetAccountId()
	if account == 0 {
		s.logger.Trace().Msg("empty account value")
		return &kvstorev1.GetKeyResponse{}, kv.ErrInvalidAccount
	}

	bucketName := request.GetBucketName()
	if bucketName == "" {
		s.logger.Trace().Msg("empty bucket name")
		return &kvstorev1.GetKeyResponse{}, kv.ErrInvalidOwner
	}

	keyName := request.GetKey()
	if keyName == "" {
		s.logger.Trace().Msg("empty key name")
		return &kvstorev1.GetKeyResponse{}, errors.New("empty key name")
	}

	req := &kvstorev1.KVStoreWrapper{
		Account: request.GetAccountId(),
		Bucket:  bucketName,
		Typ:     kvstorev1.KVStoreWrapper_REQUEST_TYPE_GET_KEY_REQUEST,
		Payload: &kvstorev1.KVStoreWrapper_GetKeyRequest{
			GetKeyRequest: request,
		},
	}

	cmd, err := req.MarshalVT()
	if err != nil {
		s.logger.Error().Err(err).Msg("can't marshal request")
		return &kvstorev1.GetKeyResponse{}, errors.Wrap(err, "can't marshal request")
	}

	ctx, cancel := context.WithTimeout(context.Background(), utils.Timeout(s.defaultTimeout))
	defer cancel()

	shard := s.shardRouter.AccountToShard(account)

	result, err := s.nh.SyncRead(ctx, shard, cmd)
	if err != nil {
		s.logger.Error().Err(err).Msg("can't apply message")
		return &kvstorev1.GetKeyResponse{}, errors.Wrap(err, "can't apply message")
	}

	resp := &kvstorev1.KVStoreWrapper{}
	err = resp.UnmarshalVT(result.([]byte))
	if err != nil {
		s.logger.Error().Err(err).Msg("can't unmarshal response")
		return &kvstorev1.GetKeyResponse{}, errors.Wrap(err, "can't unmarshal response")
	}

	if resp.Typ == kvstorev1.KVStoreWrapper_REQUEST_TYPE_RECOVERABLE_ERROR {
		errMsg := &aerrs.Error{}
		errMsg = resp.GetError()

		return &kvstorev1.GetKeyResponse{}, errors.New(errMsg.Message)
	}

	response := &kvstorev1.GetKeyResponse{}
	kvp := resp.GetGetKeyReply()
	if kvp == nil {
		response.KeyValuePair = &kvstorev1.KeyValue{}
	} else {
		response.KeyValuePair = kvp.GetKeyValuePair()
	}

	return response, nil
}

func (s *bboltStoreManager) PutKey(request *kvstorev1.PutKeyRequest) (*kvstorev1.PutKeyResponse, error) {
	account := request.GetAccountId()
	if account == 0 {
		s.logger.Trace().Msg("empty account value")
		return &kvstorev1.PutKeyResponse{}, kv.ErrInvalidAccount
	}

	bucketName := request.GetBucketName()
	if bucketName == "" {
		s.logger.Trace().Msg("empty bucketName value")
		return &kvstorev1.PutKeyResponse{}, kv.ErrInvalidOwner
	}

	keyValuePair := request.GetKeyValuePair()
	if keyValuePair == nil {
		s.logger.Trace().Msg("empty key value pair")
		return &kvstorev1.PutKeyResponse{}, kv.ErrInvalidBucketName
	}

	req := &kvstorev1.KVStoreWrapper{
		Account: request.GetAccountId(),
		Bucket:  bucketName,
		Typ:     kvstorev1.KVStoreWrapper_REQUEST_TYPE_PUT_KEY_REQUEST,
		Payload: &kvstorev1.KVStoreWrapper_PutKeyRequest{
			PutKeyRequest: request,
		},
	}

	cmd, err := req.MarshalVT()
	if err != nil {
		s.logger.Error().Err(err).Msg("can't marshal request")
		return &kvstorev1.PutKeyResponse{}, errors.Wrap(err, "can't marshal request")
	}

	ctx, cancel := context.WithTimeout(context.Background(), utils.Timeout(s.defaultTimeout))
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
		return &kvstorev1.PutKeyResponse{}, errors.Wrap(err, "can't apply message")
	}

	if cs.SeriesID != client.NoOPSeriesID {
		cs.ProposalCompleted()
	}

	resp := &kvstorev1.KVStoreWrapper{}
	err = resp.UnmarshalVT(result.Data)
	if err != nil {
		s.logger.Error().Err(err).Msg("can't unmarshal response")
		return &kvstorev1.PutKeyResponse{}, errors.Wrap(err, "can't unmarshal response")
	}

	if resp.Typ == kvstorev1.KVStoreWrapper_REQUEST_TYPE_RECOVERABLE_ERROR {
		errMsg := &aerrs.Error{}
		errMsg = resp.GetError()

		return &kvstorev1.PutKeyResponse{}, errors.New(errMsg.Message)
	}

	response := &kvstorev1.PutKeyResponse{}

	if request.Transaction == nil || cs.SeriesID != client.NoOPSeriesID {
		response.Transaction = &kvstorev1.Transaction{}
	} else {
		response.Transaction = csToTransaction(*cs)
	}

	return response, nil
}

func (s *bboltStoreManager) DeleteKey(request *kvstorev1.DeleteKeyRequest) (*kvstorev1.DeleteKeyResponse, error) {
	account := request.GetAccountId()
	if account == 0 {
		s.logger.Trace().Msg("empty account value")
		return &kvstorev1.DeleteKeyResponse{}, kv.ErrInvalidAccount
	}

	bucketName := request.GetBucketName()
	if bucketName == "" {
		s.logger.Trace().Msg("empty bucket name")
		return &kvstorev1.DeleteKeyResponse{}, kv.ErrInvalidOwner
	}

	key := request.GetKey()
	if key == "" {
		s.logger.Trace().Msg("empty key value pair")
		return &kvstorev1.DeleteKeyResponse{}, kv.ErrInvalidBucketName
	}

	req := &kvstorev1.KVStoreWrapper{
		Account: request.GetAccountId(),
		Bucket:  bucketName,
		Typ:     kvstorev1.KVStoreWrapper_REQUEST_TYPE_DELETE_KEY_REQUEST,
		Payload: &kvstorev1.KVStoreWrapper_DeleteKeyRequest{
			DeleteKeyRequest: request,
		},
	}

	cmd, err := req.MarshalVT()
	if err != nil {
		s.logger.Error().Err(err).Msg("can't marshal request")
		return &kvstorev1.DeleteKeyResponse{}, errors.Wrap(err, "can't marshal request")
	}

	ctx, cancel := context.WithTimeout(context.Background(), utils.Timeout(s.defaultTimeout))
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
		return &kvstorev1.DeleteKeyResponse{}, errors.Wrap(err, "can't apply message")
	}

	if cs.SeriesID != client.NoOPSeriesID {
		cs.ProposalCompleted()
	}

	resp := &kvstorev1.KVStoreWrapper{}
	err = resp.UnmarshalVT(result.Data)
	if err != nil {
		s.logger.Error().Err(err).Msg("can't unmarshal response")
		return &kvstorev1.DeleteKeyResponse{}, errors.Wrap(err, "can't unmarshal response")
	}

	if resp.Typ == kvstorev1.KVStoreWrapper_REQUEST_TYPE_RECOVERABLE_ERROR {
		errMsg := &aerrs.Error{}
		errMsg = resp.GetError()

		return &kvstorev1.DeleteKeyResponse{}, errors.New(errMsg.Message)
	}

	response := &kvstorev1.DeleteKeyResponse{}

	if request.Transaction == nil || cs.SeriesID != client.NoOPSeriesID {
		response.Transaction = &kvstorev1.Transaction{}
	} else {
		response.Transaction = csToTransaction(*cs)
	}

	if resp.GetDeleteKeyReply() != nil {
		response.Ok = resp.GetDeleteKeyReply().GetOk()
	} else {
		response.Ok = false
	}

	return response, nil
}
