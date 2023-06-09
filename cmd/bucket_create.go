/*
 * Copyright (c) 2022-2023 Sienna Lloyd
 *
 * Licensed under the PolyForm Internal Use License 1.0.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License here:
 *  https://github.com/mxplusb/pleiades/blob/mainline/LICENSE
 */

package cmd

import (
	"context"
	"fmt"

	kvstorev1 "github.com/mxplusb/pleiades/pkg/api/kvstore/v1"
	"github.com/mxplusb/pleiades/pkg/api/kvstore/v1/kvstorev1connect"
	"github.com/bufbuild/connect-go"
	"github.com/mitchellh/cli"
	"github.com/posener/complete"
)

var (
	_ cli.Command             = (*BucketCreateCommand)(nil)
	_ cli.CommandAutocomplete = (*BucketCreateCommand)(nil)
)

type BucketCreateCommand struct {
	*BaseCommand

	flagBucketOwner string
	flagAccountId   uint64
	flagBucketName  string
}

func (a *BucketCreateCommand) Flags() *FlagSets {
	set := a.flagSet(FlagSetHTTP | FlagSetFormat | FlagSetLogging)
	f := set.NewFlagSet("Account Options")

	f.Uint64Var(&Uint64Var{
		Name:              "account-id",
		Usage:             "Account ID to associate the bucket with.",
		Target:            &a.flagAccountId,
		Completion:        complete.PredictNothing,
		ConfigurationPath: "client.bucket.create.account-id",
	})

	f.StringVar(&StringVar{
		Name:              "bucket-owner-email",
		Usage:             "Email address of the bucket owner.",
		Target:            &a.flagBucketOwner,
		Completion:        complete.PredictNothing,
		ConfigurationPath: "client.bucket.create.account-owner-email",
	})

	f.StringVar(&StringVar{
		Name:              "bucket",
		Usage:             "Name of the bucket.",
		Target:            &a.flagBucketName,
		Completion:        complete.PredictNothing,
		ConfigurationPath: "client.bucket.create.name",
	})

	return set
}

func (a *BucketCreateCommand) AutocompleteArgs() complete.Predictor {
	return complete.PredictNothing
}

func (a *BucketCreateCommand) AutocompleteFlags() complete.Flags {
	return a.Flags().Completions()
}

func (a *BucketCreateCommand) Help() string {
	helpText := `Create a bucket.

` + a.Flags().Help()

	return helpText
}

func (a *BucketCreateCommand) Run(args []string) int {
	f := a.Flags()

	if err := f.Parse(args); err != nil {
		a.UI.Error(err.Error())
		return exitCodeFailureToParseArgs
	}

	trace := config.GetBool("logging.trace")
	if trace {
		OutputData(a.UI, config.AllSettings())
	}

	if a.flagAccountId == 0 {
		a.UI.Error("the account-id must not be 0")
		return exitCodeGenericBad
	}

	httpClient, err := a.Client()
	if err != nil {
		a.UI.Error(err.Error())
		return exitCodeGenericBad
	}

	client := kvstorev1connect.NewKvStoreServiceClient(httpClient, a.BaseCommand.flagHost)

	descriptor, err := client.CreateBucket(context.Background(), connect.NewRequest(&kvstorev1.CreateBucketRequest{
		AccountId:   a.flagAccountId,
		Owner:       a.flagBucketOwner,
		Name:        a.flagBucketName,
		Transaction: nil,
	}))
	if err != nil {
		a.UI.Error(fmt.Sprintf("error creating bucket: %s", err))
		return exitCodeRemote
	}

	if descriptor.Msg != nil {
		OutputData(a.UI, descriptor.Msg)
	}

	return exitCodeGood
}

func (a *BucketCreateCommand) Synopsis() string {
	return "create an account in the key value store"
}
