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
	"testing"

	"github.com/heimdalr/dag"
	"github.com/stretchr/testify/suite"
)

func TestCtorTestSuite(t *testing.T) {
	suite.Run(t, new(ctorTestSuite))
}

type ctorTestSuite struct {
	suite.Suite
}

var (
	_ dag.IDInterface = (*thing)(nil)
	_ dag.IDInterface = (*otherThing)(nil)

	newThingPkg = "servicegraph.newThing"
	thingType = "servicegraph.thing"
	newthingWithFooPkg = "servicegraph.newThingWithFoo"
)

type thing struct {
	foo string
}

func (t thing) ID() string {
	return "thing"
}

func newThing() *thing {
	return &thing{}
}

func newThingWithFoo(foo string) *thing {
	return &thing{
		foo: foo,
	}
}

type otherThing struct {
	bar string
	thing *thing
}

func (t otherThing) ID() string {
	return "otherThing"
}

func newOtherThing() *otherThing {
	return &otherThing{}
}

func newOtherThingWithBar(bar string) *otherThing {
	return &otherThing{
		bar: bar,
	}
}

func newOtherThingWithDependencyOnThing(thing *thing) *otherThing {
	return &otherThing{
		bar: thing.foo,
		thing: thing,
	}
}

type transitiveDependency struct {
	thing *thing
}

func newTransitiveDependency(otherThing *otherThing) *transitiveDependency {
	return &transitiveDependency{
		thing: otherThing.thing,
	}
}

func newTransitiveDependencyWhichProvidesThing(otherThing *otherThing) *thing {
	return &thing{
		foo: otherThing.bar,
	}
}

func (t *ctorTestSuite) TestNewCtorNode() {
	d := dag.NewDAG()

	// Test with a function that returns a pointer to a struct
	node, err := newCtorNode(newThing, d)
	t.Require().NoError(err, "failed to create ctorNode")
	t.Require().NotNil(node, "ctorNode is nil")

	newThingNode, err := d.GetVertex(newThingPkg)
	t.Require().NoError(err, "failed to get vertex")
	t.Require().NotNil(newThingNode, "vertex is nil")

	newThingNodeCast, ok := newThingNode.(*ctorNode)
	t.Require().True(ok, "failed to cast node to ctorNode")
	t.Require().Equal(newThingPkg, newThingNodeCast.name, "ctorNode name is incorrect")

	thingNode, err := d.GetVertex(thingType)
	t.Require().NoError(err, "failed to get vertex")
	t.Require().NotNil(thingNode, "vertex is nil")

	// ensure the node is of type *typeNode
	tnCast, ok := thingNode.(*typeNode)
	t.Require().True(ok, "failed to cast node to thing")
	t.Require().IsType(&typeNode{}, tnCast, "thing is not of type *thing")

	// ensure the edge exists with dag.IsEdge
	ok, err = d.IsEdge(newThingNodeCast.ID(), tnCast.ID())
	t.Require().NoError(err, "failed to check for edge")
	t.Require().True(ok, "edge does not exist")

	// ensure duplicate nodes are not created
	node, err = newCtorNode(newThing, d)
	t.Require().Error(err, "duplicate ctorNode created")
	t.Require().Nil(node, "ctorNode is not nil")

	t.T().Logf("graph: %s", d.String())
}

func (t *ctorTestSuite) TestNewCtorNodeWithBasicArgument() {
	// Test with a function that returns a pointer to a struct and takes an argument
	d := dag.NewDAG()

	node, err := newCtorNode(newThingWithFoo, d)
	t.Require().NoError(err, "failed to create ctorNode")
	t.Require().NotNil(node, "ctorNode is nil")

	newThingWithFooNode, err := d.GetVertex(newthingWithFooPkg)
	t.Require().NoError(err, "failed to get vertex")
	t.Require().NotNil(newThingWithFooNode, "vertex is nil")

	thingNode, err := d.GetVertex(thingType)
	t.Require().NoError(err, "failed to get vertex")
	t.Require().NotNil(thingNode, "vertex is nil")

	// ensure the node is of type *typeNode
	tnCast, ok := thingNode.(*typeNode)
	t.Require().True(ok, "failed to cast node to thing")
	t.Require().IsType(&typeNode{}, tnCast, "thing is not of type *thing")

	t.T().Logf("graph: %s", d.String())
}

func (t *ctorTestSuite) TestNewCtorWithBasicDependencies() {
	d := dag.NewDAG()

	//newOtherThingPkg := "servicegraph.newOtherThing"
	//otherThingPkg := "servicegraph.otherThing"
	//newOtherThingWithBarPkg := "servicegraph.newOtherThingWithBar"
	//newOtherThingWithDependencyOnThingPkg := "servicegraph.newOtherThingWithDependencyOnThing"

	otherThingNode, err := newCtorNode(newOtherThing, d)
	t.Require().NoError(err, "failed to create ctorNode")
	t.Require().NotNil(otherThingNode, "ctorNode is nil")

	otherThingWithDependencyNode, err := newCtorNode(newOtherThingWithDependencyOnThing, d)
	t.Require().NoError(err, "failed to create ctorNode")
	t.Require().NotNil(otherThingWithDependencyNode, "ctorNode is nil")

	// add a new node which already has a dependency available
	thingNode, err := newCtorNode(newThing, d)
	t.Require().NoError(err, "failed to create ctorNode")
	t.Require().NotNil(thingNode, "ctorNode is nil")

	transitiveDependencyNode, err := newCtorNode(newTransitiveDependency, d)
	t.Require().NoError(err, "failed to create ctorNode")
	t.Require().NotNil(transitiveDependencyNode, "ctorNode is nil")

	// create a dependency loop
	transitiveDependencyWhichProvidesThingNode, err := newCtorNode(newTransitiveDependencyWhichProvidesThing, d)
	t.Require().Error(err, "failed to create ctorNode")
	t.Require().Nil(transitiveDependencyWhichProvidesThingNode, "ctorNode is nil")

	t.T().Logf("graph: %s", d.String())
}