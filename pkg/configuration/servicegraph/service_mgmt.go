/*
 * Copyright (c) 2023 Sienna Lloyd
 *
 * Licensed under the PolyForm Internal Use License 1.0.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License here:
 *  https://github.com/mxplusb/pleiades/blob/mainline/LICENSE
 */

package servicegraph

import (
	"context"
	"reflect"

	"github.com/cockroachdb/errors"
	"github.com/heimdalr/dag"
	zlog "github.com/rs/zerolog"
)

type LifecycleState int
const (
	Stopped LifecycleState = iota
	Starting
	Started
	Stopping
)

type IHostedService interface {
	dag.IDInterface
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	LifecycleState() LifecycleState
}

type IOption interface {
	dag.IDInterface
}

type ServiceManager struct {
	logger      zlog.Logger
	dag         *dag.DAG
}

// Create a new ServiceManager
func NewServiceManager(logger zlog.Logger) *ServiceManager {
	return &ServiceManager{logger: logger.With().Str("aspect", "service-manager").Logger(), dag: dag.NewDAG()}
}

// Boot will bootstrap the entire service graph in the proper order
func (s *ServiceManager) Boot() error {
	return nil
}

func (s *ServiceManager) Shutdown() error {
	return nil
}

func (s *ServiceManager) AddOption(opt IOption) error {
	t := reflect.TypeOf(opt).Kind()
	if t != reflect.Struct {
		return errors.Newf("can't add an option that's not a struct")
	}

	_, err := s.dag.AddVertex(t)
	if err != nil {
		return errors.Wrapf(err, "can't add options to service graph")
	}

	return nil
}

var _ dag.Visitor = (*walker)(nil)
type walker struct {
	targets *[]string
}

func (w walker) Visit(vertexer dag.Vertexer) {

}
