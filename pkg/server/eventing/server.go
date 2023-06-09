/*
 * Copyright (c) 2022-2023 Sienna Lloyd
 *
 * Licensed under the PolyForm Internal Use License 1.0.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License here:
 *  https://github.com/mxplusb/pleiades/blob/mainline/LICENSE
 */

package eventing

import (
	"context"

	"github.com/mxplusb/pleiades/pkg/messaging"
	"github.com/rs/zerolog"
	"go.uber.org/fx"
)

var (
	serverSingleton *Server
)

type EventServerBuilderParams struct {
	fx.In
	Logger zerolog.Logger
}

type EventServerBuilderResults struct {
	fx.Out
	Server *Server
}

func NewServer(lc fx.Lifecycle, params EventServerBuilderParams) EventServerBuilderResults {
	srv, err := messaging.NewEmbeddedMessagingWithDefaults(params.Logger)
	if err != nil {
		params.Logger.Fatal().Err(err).Msg("can't create embedded message bus")
	}

	serverSingleton = &Server{srv, params.Logger.With().Str("component", "eventing").Logger()}

	// this is started so the other constructors register properly
	serverSingleton.Start()

	// lifecycle hooks
	lc.Append(fx.Hook{
		// empty hook for fx
		OnStart: func(_ context.Context) error {
			return nil
		},
		OnStop: func(ctx context.Context) error {
			serverSingleton.Stop()
			return nil
		}})

	return EventServerBuilderResults{
		Server: serverSingleton,
	}
}

type Server struct {
	*messaging.EmbeddedMessaging
	logger zerolog.Logger
}

func (s *Server) GetRaftEventHandler() (*messaging.RaftEventHandler, error) {
	pubSubClient, err := s.EmbeddedMessaging.GetPubSubClient()
	if err != nil {
		s.logger.Error().Err(err).Msg("can't create pubsub client")
		return nil, err
	}

	queueClient, err := s.EmbeddedMessaging.GetStreamClient()
	if err != nil {
		s.logger.Error().Err(err).Msg("can't create stream client")
		return nil, err
	}

	return messaging.NewRaftEventHandler(pubSubClient, queueClient, s.logger), nil
}

func (s *Server) GetRaftSystemEventListener() (*messaging.RaftSystemListener, error) {
	pubSubClient, err := s.EmbeddedMessaging.GetPubSubClient()
	if err != nil {
		s.logger.Error().Err(err).Msg("can't create pubsub client")
		return nil, err
	}

	queueClient, err := s.EmbeddedMessaging.GetStreamClient()
	if err != nil {
		s.logger.Error().Err(err).Msg("can't create stream client")
		return nil, err
	}

	return messaging.NewRaftSystemListener(pubSubClient, queueClient, s.logger)
}

type NewPubSubClientBuilderParams struct {
	fx.In

	Server *Server
}

func NewPubSubClient(params NewPubSubClientBuilderParams) (*messaging.EmbeddedMessagingPubSubClient, error) {
	return params.Server.GetPubSubClient()
}

type NewStreamClientBuilderParams struct {
	fx.In

	Server *Server
}

func NewStreamClient(params NewStreamClientBuilderParams) (*messaging.EmbeddedMessagingStreamClient, error) {
	return params.Server.GetStreamClient()
}