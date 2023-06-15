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
	case t.Kind() == reflect.Struct && t.Kind() != reflect.Ptr:
		return &typeNode{}, errors.Newf("param %v must be a pointer because it is a struct", t)
	}

	tn := &typeNode{}

	if t.Kind() == reflect.Ptr {
		tn.optional = true
		tn.name = t.Elem().Name()
		tn.typ = t.Elem()
	} else {
		tn.optional = false
		tn.name = t.Name()
		tn.typ = t
	}

	if err := d.AddVertexByID(tn.ID(), tn); err != nil {
		// it already exists so we can skip it
		if strings.Contains(err.Error(), "already known") {
			return tn, nil
		}
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
	return p.typ.String()
}

// String returns a string representation of the typeNode
func (p typeNode) String() string {
	var opts []string
	if p.optional {
		opts = append(opts, "optional")
	}

	if len(opts) == 0 {
		return fmt.Sprint(p.typ)
	}

	// [optional, name="foo"]
	return fmt.Sprintf("%v[%v]", p.typ, strings.Join(opts, ", "))
}
