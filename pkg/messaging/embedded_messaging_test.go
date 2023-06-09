/*
 * Copyright (c) 2022-2023 Sienna Lloyd
 *
 * Licensed under the PolyForm Internal Use License 1.0.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License here:
 *  https://github.com/mxplusb/pleiades/blob/mainline/LICENSE
 */

package messaging

import (
	"testing"

	"github.com/mxplusb/pleiades/pkg/utils"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/suite"
)

func TestEmbeddedEventStream(t *testing.T) {
	suite.Run(t, new(EmbeddedEventStreamTestSuite))
}

type EmbeddedEventStreamTestSuite struct {
	suite.Suite
	opts   *EmbeddedMessagingStreamOpts
	logger zerolog.Logger
}

func (t *EmbeddedEventStreamTestSuite) SetupSuite() {
	t.logger = utils.NewTestLogger(t.T())
}

func (t *EmbeddedEventStreamTestSuite) TearDownTest() {
	singleton.Stop()
	singleton = nil
}

func (t *EmbeddedEventStreamTestSuite) TestNew() {
	e, err := NewEmbeddedMessagingWithDefaults(t.logger)
	t.Require().NoError(err, "there must not be an error creating a new embedded event stream")
	t.Require().NotNil(e, "the event stream must not be nil")
}

func (t *EmbeddedEventStreamTestSuite) TestStartAndStop() {
	e, err := NewEmbeddedMessagingWithDefaults(t.logger)
	t.Require().NoError(err, "there must not be an error creating a new embedded event stream")
	t.Require().NotNil(e, "the event stream must not be nil")

	t.Require().NotPanics(e.Start, "the embedded server must not panic on start")

	t.Require().NotPanics(e.Stop, "the embedded server must not panic on stop")
}

func (t *EmbeddedEventStreamTestSuite) TestGetPubSubClient() {
	e, err := NewEmbeddedMessagingWithDefaults(t.logger)
	t.Require().NoError(err, "there must not be an error creating a new embedded event stream")
	t.Require().NotNil(e, "the event stream must not be nil")

	t.Require().NotPanics(func() {
		e.Start()
	}, "the embedded server must not panic on start")

	embeddedClient, err := e.GetPubSubClient()
	t.Require().NoError(err, "there must not be an error when creating an embedded pubSubClient")
	t.Require().NotNil(embeddedClient, "the pubSubClient must not be nil")
}

func (t *EmbeddedEventStreamTestSuite) TestGetStreamClient() {
	e, err := NewEmbeddedMessagingWithDefaults(t.logger)
	t.Require().NoError(err, "there must not be an error creating a new embedded event stream")
	t.Require().NotNil(e, "the event stream must not be nil")

	t.Require().NotPanics(func() {
		e.Start()
	}, "the embedded server must not panic on start")

	embeddedClient, err := e.GetStreamClient()
	t.Require().NoError(err, "there must not be an error when creating an embedded pubSubClient")
	t.Require().NotNil(embeddedClient, "the pubSubClient must not be nil")
}
