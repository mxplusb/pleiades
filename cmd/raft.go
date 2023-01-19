/*
 * Copyright (c) 2022-2023 Sienna Lloyd
 *
 * Licensed under the PolyForm Strict License 1.0.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License here:
 *  https://github.com/mxplusb/pleiades/blob/mainline/LICENSE
 */
package cmd

import (
	"github.com/spf13/cobra"
)

// raftCmd represents the raft command
var raftCmd = &cobra.Command{
	Use:   "raft",
	Short: "operations on the raft subsystem",
	Long: `various operations used to control various aspects of the raft subsystems`,
}

func init() {
	rootCmd.AddCommand(raftCmd)

	raftCmd.PersistentFlags().String("host", "http://localhost:8080", "target host for a pleiades cluster")
	//config.BindPFlag("client.grpcAddr", raftCmd.PersistentFlags().Lookup("host"))
}