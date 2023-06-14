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
	"reflect"

	"github.com/cockroachdb/errors"
	"github.com/heimdalr/dag"
)

var (
	_ dag.IDInterface = (*ctorNode)(nil)
)

func newCtorNode(ctor interface{}, d *dag.DAG) (*ctorNode, error) {
	cval := reflect.ValueOf(ctor)
	ctype := cval.Type()

	c := &ctorNode{
		name: cval.String(),
		called: false,
	}
	c.def = inspectFunc(ctor)

	// add the ctor to the dag
	if err:= d.AddVertexByID(c.ID(), c); err != nil {
		return nil, errors.Wrapf(err, "failed to add ctor %v to service graph", c)
	}

	// add incoming parameters to the dag
	if ctype.NumIn() > 0 {
		for i := 0; i < ctype.NumIn(); i++ {
			switch {
			case ctype.Kind() == reflect.Func:
				node, err := newCtorNode(ctype.In(i), d)
				if err != nil {
					return nil, errors.Wrapf(err, "failed to determine ctor node for ctor %v param %v", ctype, i)
				}
				if err := d.AddEdge(node.ID(), c.ID()); err != nil {
					return nil, errors.Wrapf(err, "failed to add param dependency from %v to %v", node, c)
				}
			default:
				in := ctype.In(i)
				inNode, err := newTypeNode(in, d)
				if err != nil {
					return nil, errors.Wrapf(err, "failed to determine type node for ctor %v param %v", ctype, i)
				}

				if err := d.AddEdge(inNode.ID(), c.ID()); err != nil {
					return nil, errors.Wrapf(err, "failed to add dependency from %v to %v", inNode, c)
				}
			}
		}
	}

	// add outgoing results to the dag
	if ctype.NumOut() > 0 {
		for i := 0; i < ctype.NumOut(); i++ {
			switch {
			case ctype.Kind() == reflect.Func:
				node, err := newCtorNode(ctype.Out(i), d)
				if err != nil {
					return nil, errors.Wrapf(err, "failed to determine ctor node for ctor %v param %v", ctype, i)
				}
				if err := d.AddEdge(c.ID(), node.ID()); err != nil {
					return nil, errors.Wrapf(err, "failed to add param dependency from %v to %v", node, c)
				}
			default:
				out := ctype.Out(i)
				outNode, err := newTypeNode(out, d)
				if err != nil {
					return nil, errors.Wrapf(err, "failed to determine type node for ctor %v param %v", ctype, i)
				}

				if err := d.AddEdge(c.ID(), outNode.ID()); err != nil {
					return nil, errors.Wrapf(err, "failed to add dependency from %v to %v", outNode, c)
				}
			}
		}
	}

	return c, nil
}

type ctorNode struct {
	// ctor metadata
	ctor interface{}
	ctype reflect.Type

	// location of the ctor definition
	def *funcType

	name string
	called bool
}

func (c *ctorNode) ID() string {
	return c.name
}
