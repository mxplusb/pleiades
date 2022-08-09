/*
 * Copyright (c) 2022 Sienna Lloyd
 *
 * Licensed under the PolyForm Strict License 1.0.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License here:
 *  https://github.com/mxplusb/pleiades/blob/mainline/LICENSE
 */

package blaze

import (
	"github.com/mxplusb/pleiades/api/v1/database"
	"github.com/mxplusb/pleiades/pkg/conf"
	"github.com/aperturerobotics/starpc/srpc"
	dconfig "github.com/lni/dragonboat/v3/config"
)

func RegisterNodeHostRpcServer(mux srpc.Mux, conf *conf.NodeHostConfig, clogger conf.Logger) error {
	logger := clogger.GetLogger()
	l := logger.With().Str("service", "node-host").Logger()

	translatedConf, err := buildNodeHostConfig(*conf)
	if err != nil {
		l.Error().Err(err).Msg("failed to build node host config")
		return err
	}

	node, err := NewNode(translatedConf, clogger)
	if err != nil {
		l.Error().Err(err).Msg("failed to build node")
		return err
	}

	nodeHostRpc := NewNodeHostRPCServer(node, l)

	if err := database.SRPCRegisterRaftControlService(mux, nodeHostRpc); err != nil {
		l.Error().Err(err).Msg("failed to register raft control service")
		return err
	}

	return nil
}

func buildNodeHostConfig(conf conf.NodeHostConfig) (dconfig.NodeHostConfig, error) {
	dconf := dconfig.NodeHostConfig{
		DeploymentID: conf.DeploymentID,
		WALDir: conf.WALDir,
		NodeHostDir: conf.NodeHostDir,
		RTTMillisecond: conf.RTTMillisecond,
		RaftAddress: conf.RaftAddress,
		AddressByNodeHostID: conf.AddressByNodeHostID,
		ListenAddress: conf.ListenAddress,
		MutualTLS: conf.MutualTLS,
		CAFile: conf.CAFile,
		CertFile: conf.CertFile,
		KeyFile: conf.KeyFile,
		EnableMetrics: conf.EnableMetrics,
		MaxSnapshotRecvBytesPerSecond: conf.MaxSnapshotRecvBytesPerSecond,
		MaxSnapshotSendBytesPerSecond: conf.MaxSnapshotSendBytesPerSecond,
		MaxReceiveQueueSize: conf.MaxReceiveQueueSize,
		MaxSendQueueSize: conf.MaxSendQueueSize,
		NotifyCommit: conf.NotifyCommit,
		Gossip: dconfig.GossipConfig{
			BindAddress:      conf.Gossip.BindAddress,
			AdvertiseAddress: conf.Gossip.AdvertiseAddress,
			Seed:             conf.Gossip.Seed,
		},
	}
	return dconf, nil
}
