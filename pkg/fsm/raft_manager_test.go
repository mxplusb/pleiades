package fsm

import (
	"os"
	"testing"

	"github.com/hashicorp/consul/api"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.uber.org/fx/fxtest"
	"r3t.io/pleiades/pkg/conf"
	configv1 "r3t.io/pleiades/pkg/pb/config/v1"
	"r3t.io/pleiades/pkg/servers/services"
	"r3t.io/pleiades/pkg/utils"
)

type RaftManagerTests struct {
	suite.Suite
	logger    zerolog.Logger
	store     *services.StoreManager
	lifecycle *fxtest.Lifecycle
	client    *api.Client
	env       *conf.EnvironmentConfig
}

func TestRaftManager(t *testing.T) {
	suite.Run(t, new(RaftManagerTests))
}

func (rmt *RaftManagerTests) SetupSuite() {
	rmt.logger = utils.NewTestLogger(rmt.T())

	var err error
	rmt.lifecycle = fxtest.NewLifecycle(rmt.T())
	rmt.client, err = conf.NewConsulClient(rmt.lifecycle)
	require.Nil(rmt.T(), err, "failed to connect to consul")
	require.NotNil(rmt.T(), rmt.client, "the consul client can't be empty")

	rmt.env, err = conf.NewEnvironmentConfig(rmt.client)
	require.Nil(rmt.T(), err, "the environment config is needed")
	require.NotNil(rmt.T(), rmt.env, "the environment config must be rendered")

	rmt.store = services.NewStoreManager(rmt.T().TempDir(), rmt.logger)
	err = rmt.store.Start(false)
	require.Nil(rmt.T(), err, "there can't be an error when starting the store manager")
}

func (rmt *RaftManagerTests) BeforeTest(suiteName, testName string) {
	err := os.RemoveAll("tmp")
	require.Nil(rmt.T(), err, "there must not be an error clearing out the temporary test directory")
}

func (rmt *RaftManagerTests) TestNewRaftManager() {
	require.NotPanics(rmt.T(), func() {
		NewOldRaftManager(rmt.store, rmt.logger)
	}, "building a new raft manager should not panic")
}

func (rmt *RaftManagerTests) TestRaftManagerPut() {
	var manager *OldRaftManager[configv1.RaftConfig]
	require.NotPanics(rmt.T(), func() {
		manager = NewOldRaftManager(rmt.store, rmt.logger)
	}, "building a new raft manager should not panic")

	err := manager.Put("test", &configv1.RaftConfig{ClusterId: 123})
	require.Nil(rmt.T(), err, "there must not be an error when storing a raft config")
}

func (rmt *RaftManagerTests) TestRaftManagerPutAndGet() {
	var manager *OldRaftManager[configv1.RaftConfig]
	require.NotPanics(rmt.T(), func() {
		manager = NewOldRaftManager(rmt.store, rmt.logger)
	}, "building a new raft manager should not panic")

	testStruct := &configv1.RaftConfig{ClusterId: 123}
	err := manager.Put("test", testStruct)
	require.Nil(rmt.T(), err, "there must not be an error when storing a raft config")

	resp, err := manager.Get("test")
	if err != nil {
		require.Nil(rmt.T(), err, "there must not be a fetch error")
	}
	assert.Equal(rmt.T(), testStruct.ClusterId, resp.ClusterId, "the fetched type must match the stored type")

	testStruct.ClusterId = 456
	err = manager.Put("another-test", testStruct)
	require.Nil(rmt.T(), err, "there must not be an error when storing a raft config")

	resp, err = manager.Get("another-test")
	require.Nil(rmt.T(), err, "there must not be an error when storing a second raft config")

	resp, err = manager.Get("another-test")
	if err != nil {
		require.Nil(rmt.T(), err, "there must not be a fetch error")
	}
	assert.Equal(rmt.T(), testStruct.ClusterId, resp.ClusterId, "the fetched type must match the second stored type")
}

func (rmt *RaftManagerTests) TestRaftManagerGetAll() {
	var manager *OldRaftManager[configv1.RaftConfig]
	require.NotPanics(rmt.T(), func() {
		manager = NewOldRaftManager(rmt.store, rmt.logger)
	}, "building a new raft manager should not panic")

	testStructOne := &configv1.RaftConfig{ClusterId: 123, Id: "testStructOne"}
	err := manager.Put("test", testStructOne)
	require.Nil(rmt.T(), err, "there must not be an error when storing a raft config")

	testStructTwo := &configv1.RaftConfig{ClusterId: 456, Id: "testStructTwo"}
	err = manager.Put("another-test", testStructTwo)
	require.Nil(rmt.T(), err, "there must not be an error when storing a second raft config")

	all, err := manager.GetAll()
	require.Nil(rmt.T(), err, "there must be no error when fetching all the stored raft configs")

	tests := map[string]*configv1.RaftConfig{
		"testStructOne": testStructOne,
		"testStructTwo": testStructTwo,
	}

	for idx, _ := range all {
		targetVal := all[idx]
		ref, ok := tests[targetVal.Id]
		rmt.Require().True(ok, "the target value must exist")
		rmt.Assert().Equal(ref.ClusterId, targetVal.ClusterId, "the results should be the same")
	}
}
