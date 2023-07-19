/*
 * Copyright (c) 2022-2023 Sienna Lloyd
 *
 * Licensed under the PolyForm Internal Use License 1.0.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License here:
 *  https://github.com/mxplusb/pleiades/blob/mainline/LICENSE
 */

package cli

import (
	"context"
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/mitchellh/cli"
	"github.com/mitchellh/go-wordwrap"
	"github.com/mxplusb/pleiades/pkg/raftpb"
	"github.com/mxplusb/pleiades/pkg/raftpb/raftpbconnect"
	"github.com/posener/complete"
)

var (
	_ cli.Command             = (*FabricStartReplicaCommand)(nil)
	_ cli.CommandAutocomplete = (*FabricStartReplicaCommand)(nil)
)

type FabricStartReplicaCommand struct {
	*BaseCommand

	flagShardId   uint64
	flagReplicaId uint64
	flagType      string
	flagRestart   bool
}

func (f *FabricStartReplicaCommand) Flags() *FlagSets {
	set := f.flagSet(FlagSetHTTP | FlagSetFormat | FlagSetLogging | FlagSetTimeout)
	fs := set.NewFlagSet("Fabric Options")

	fs.Uint64Var(&Uint64Var{
		Name: "shard-id",
		Usage: `The ID of the new shard. This is global to the node constellation as it increases the 
data fabric size.`,
		Target:            &f.flagShardId,
		Completion:        complete.PredictNothing,
		ConfigurationPath: "fabric.start-replica.shard-id",
	})

	fs.Uint64Var(&Uint64Var{
		Name:              "replica-id",
		Usage:             `The ID of the new replica. This is specific to each shard.`,
		Target:            &f.flagReplicaId,
		Completion:        complete.PredictNothing,
		ConfigurationPath: "fabric.start-replica.replica-id",
	})

	fs.StringVar(&StringVar{
		Name: "type",
		Usage: `The type of shard to create. See the greater help message for more information on the 
specific values.`,
		Target:            &f.flagType,
		Completion:        complete.PredictSet("kvpb"),
		ConfigurationPath: "fabric.start-replica.shard-type",
	})

	fs.BoolVar(&BoolVar{
		Name:              "restart",
		Usage:             "Restart a previously stopped replica.",
		Default:           false,
		Target:            &f.flagRestart,
		Completion:        complete.PredictNothing,
		ConfigurationPath: "fabric.start-replica.restart",
	})

	return set
}

func (f *FabricStartReplicaCommand) AutocompleteArgs() complete.Predictor {
	return complete.PredictNothing
}

func (f *FabricStartReplicaCommand) AutocompleteFlags() complete.Flags {
	return f.Flags().Completions()
}

// nb (sienna): use word wrap in the editor as this will format properly in the terminal
func (f *FabricStartReplicaCommand) Help() string {
	helpText := `Start a replica.

In order for a replica to be properly provisioned, it must be started after it's created. This command starts the specific replica on the targeted host.

` + f.Flags().Help()

	return wordwrap.WrapString(helpText, 80)
}

func (f *FabricStartReplicaCommand) Run(args []string) int {
	fs := f.Flags()

	if err := fs.Parse(args); err != nil {
		f.UI.Error(err.Error())
		return exitCodeFailureToParseArgs
	}

	trace := config.GetBool("logging.trace")
	if trace {
		OutputData(f.UI, config.AllSettings())
	}

	httpClient, err := f.Client()
	if err != nil {
		f.UI.Error(err.Error())
		return exitCodeGenericBad
	}

	expiry := time.Now().UTC().Add(time.Duration(config.GetInt32("client.timeout")) * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), expiry)
	defer cancel()

	var smType raftpb.StateMachineType
	switch f.flagType {
	case "kvpb":
		smType = raftpb.StateMachineType_STATE_MACHINE_TYPE_KV
	default:
		f.UI.Error("unsupported state machine type")
		return exitCodeGenericBad
	}

	client := raftpbconnect.NewShardServiceClient(httpClient, f.BaseCommand.flagHost)

	descriptor, err := client.StartReplica(ctx, connect.NewRequest(&raftpb.StartReplicaRequest{
		ShardId:   f.flagShardId,
		ReplicaId: f.flagReplicaId,
		Type:      smType,
		Restart:   false,
	}))
	if err != nil {
		f.UI.Error(err.Error())
		return exitCodeRemote
	}

	if descriptor != nil {
		OutputData(f.UI, descriptor.Msg)
	}

	return exitCodeGood
}

func (f *FabricStartReplicaCommand) Synopsis() string {
	return "Start a replica."
}