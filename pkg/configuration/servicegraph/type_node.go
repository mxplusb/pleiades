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
	"fmt"
	"reflect"
	"strings"

	"github.com/cockroachdb/errors"
	"github.com/heimdalr/dag"
)

// newTypeNode creates a new typeNode from a reflect.Type
func newTypeNode(t reflect.Type, d *dag.DAG) (*typeNode, error) {
	switch {
	case t.Kind() != reflect.Ptr:
		return &typeNode{}, errors.Newf("param %v must be a pointer", t)
	}

	tn := &typeNode{
		name: t.Elem().Name(),
		typ:  t.Elem(),
	}

	if err := d.AddVertexByID(tn.ID(), tn); err != nil {
		return nil, errors.Wrapf(err, "failed to add  %v", tn)
	}

	return tn, nil
}

var (
	_ dag.IDInterface = (*typeNode)(nil)
)

// typeNode is a node in the graph that represents a parameter to a constructor
type typeNode struct {
	name     string
	optional bool
	typ      reflect.Type
}

func (p typeNode) ID() string {
	return p.name
}

// String returns a string representation of the typeNode
func (p typeNode) String() string {
	var opts []string
	if p.optional {
		opts = append(opts, "optional")
	}
	if p.name != "" {
		opts = append(opts, fmt.Sprintf("name=%q", p.name))
	}

	if len(opts) == 0 {
		return fmt.Sprint(p.typ)
	}

	// [optional, name="foo"]
	return fmt.Sprintf("%v[%v]", p.typ, strings.Join(opts, ", "))
}
