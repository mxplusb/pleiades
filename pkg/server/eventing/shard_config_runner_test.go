/*
 * Copyright (c) 2022 Sienna Lloyd
 *
 * Licensed under the PolyForm Strict License 1.0.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License here:
 *  https://github.com/mxplusb/pleiades/blob/mainline/LICENSE
 */

package eventing

import (
	"testing"
	"time"

	raftv1 "github.com/mxplusb/pleiades/pkg/api/raft/v1"
	"github.com/mxplusb/pleiades/pkg/configuration"
	"github.com/mxplusb/pleiades/pkg/utils"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/suite"
)

func TestShardConfigRunner(t *testing.T) {
	suite.Run(t, new(shardConfigRunnerTestSuite))
}

type shardConfigRunnerTestSuite struct {
	suite.Suite
	logger zerolog.Logger
}

func (t *shardConfigRunnerTestSuite) SetupSuite() {
	t.logger = utils.NewTestLogger(t.T())
}

func (t *shardConfigRunnerTestSuite) TestLifecycle() {
	configuration.Get().SetDefault("server.datastore.basePath", t.T().TempDir())

	nh, conf := utils.BuildTestShard(t.T())

	runner, err := NewShardConfigRunner(nh, t.logger)
	t.Require().NoError(err, "there must not be an error when creating the shard config runner")
	t.Require().NotNil(runner, "the runner must not be nil")
	go runner.run()
	defer func() { runner.done <- struct{}{} }()

	client, err := runner.msgServ.GetStreamClient()
	t.Require().NoError(err, "there must not be an error when fetching a stream client")

	testState := &raftv1.ShardStateEvent{Event: &raftv1.ShardState{ShardId: conf.ClusterID, ConfigChangeId: 2}, Cmd: raftv1.ShardStateEvent_CMD_TYPE_PUT}
	t.logger.Printf("test shard: %d", testState.GetEvent().GetShardId())
	payload, err := testState.MarshalVT()
	t.Require().NoError(err, "there must not be an error when marshaling the test payload")

	_, err = client.Publish(ShardConfigStream, payload)
	t.Require().NoError(err, "there must not be an error when publishing the event")

	// wait for the asynchronous stuff to settle a bit
	utils.Wait(1000 * time.Millisecond)

	// check the db to ensure the value was stored.
	state, err := runner.store.Get(testState.GetEvent().GetShardId())
	t.Require().NoError(err, "there must not be an error fetching the state")
	t.Require().Equal(testState.GetEvent().GetShardId(), state.GetShardId(), "the shard ids must match")
}
