/*
 * Copyright (c) 2022 Sienna Lloyd <sienna.lloyd@hey.com>
 */

package v1

import (
	"context"
	"errors"

	"capnproto.org/go/capnp/v3"
	"github.com/rs/zerolog"
	"r3t.io/pleiades/pkg/fsm"
	configv1 "r3t.io/pleiades/pkg/protocols/config/v1"
	"r3t.io/pleiades/pkg/servers/services"
)

var (
	_ configv1.ConfigService_Server = (*ConfigServiceImpl)(nil)
)

type ConfigServiceImpl struct {
	//store       *services.StoreManager
	logger      zerolog.Logger
	raftManager *fsm.ConfigServiceStoreManager
}

// NewConfigService creates a instance of the configuration service. This is a singleton.
// The configuration service is responsible for managing all the service available on a deployed host.
func NewConfigService(store *services.StoreManager, logger zerolog.Logger) (*ConfigServiceImpl, error) {
	l := logger.With().Str("service", "config").Logger()
	manager, err := fsm.NewConfigServiceStoreManager(logger, store)
	if err != nil {
		return nil, err
	}

	return &ConfigServiceImpl{
		logger:      l,
		raftManager: manager,
	}, nil
}

func (c *ConfigServiceImpl) GetConfig(ctx context.Context, call configv1.ConfigService_getConfig) error {
	req, err := call.Args().Request()
	if err != nil {
		return err
	}

	what := req.What()
	amount := req.Amount()
	switch what {
	case configv1.GetConfigurationRequest_Type_all:
	case configv1.GetConfigurationRequest_Type_raft:
		switch amount {
		case configv1.GetConfigurationRequest_Specificity_one:
			key, err := req.Id()
			if err != nil {
				c.logger.Err(err).Str("key", key).Msg("error reading key")
				return err
			}
			return c.getRaftConfig(ctx, key, call)
		}
	}

	return nil
}

func (c *ConfigServiceImpl) getRaftConfig(ctx context.Context, key string, call configv1.ConfigService_getConfig) error {
	if key == "" {
		return errors.New("cannot request a named record without a key")
	}

	// allocate the results
	res, err := call.AllocResults()
	if err != nil {
		c.logger.Err(err).Str("key", key).Msg("error allocating results")
		return err
	}

	// acknowledge the request now that we can allocate space for the results
	call.Ack()

	// find the target
	val, err := c.raftManager.Get(key)
	if err != nil {
		c.logger.Err(err).Str("key", key).Msg("error getting config")
		return err
	}

	// prep the results
	_, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		c.logger.Err(err).Str("key", key).Msg("error creating message")
		return err
	}

	// create the root config response
	config, err := configv1.NewRootGetConfigurationResponse(seg)
	if err != nil {
		c.logger.Err(err).Str("key", key).Msg("error creating getRaftConfigurationResponse")
		return err
	}

	// create the config list
	slice, err := configv1.NewRaftConfiguration_List(seg, 1)
	if err != nil {
		c.logger.Err(err).Str("key", key).Msg("error creating raftConfiguration_List")
		return err
	}

	// set the config into the slice
	err = slice.Set(0, *val)
	if err != nil {
		c.logger.Err(err).Str("key", key).Msg("error setting raftConfiguration_List")
		return err
	}

	// set the config list into the config response
	err = config.SetRaft(slice)
	if err != nil {
		c.logger.Err(err).Str("key", key).Msg("error setting config")
		return err
	}

	// map it!
	return res.SetResponse(config)
}

