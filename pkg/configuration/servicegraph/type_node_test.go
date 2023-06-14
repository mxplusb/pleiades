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
	"testing"

	"github.com/heimdalr/dag"
	"github.com/stretchr/testify/suite"
)

type typeNodeTestSuite struct {
	suite.Suite
}

func TestTypeNodeTestSuite(t *testing.T) {
	suite.Run(t, new(typeNodeTestSuite))
}

func (t *typeNodeTestSuite) TestNewTypeNode() {
	d := dag.NewDAG()

	tStructName := "testStruct"
	type testStruct struct {
		foo string
	}

	tStruct := &testStruct{
		foo: "bar",
	}

	tn, err := newTypeNode(reflect.TypeOf(tStruct), d)
	t.Require().NoError(err, "failed to create typeNode")
	t.Require().NotNil(tn, "typeNode is nil")
	t.Require().Equal(tStructName, tn.name, "typeNode name is incorrect")
	t.Require().IsType(reflect.TypeOf(tStruct), tn.typ, "typeNode type is incorrect")

	node, err := d.GetVertex(tStructName)
	t.Require().NoError(err, "failed to get vertex")
	t.Require().NotNil(node, "vertex is nil")

	tn2, ok := node.(*typeNode)
	t.Require().True(ok, "failed to cast node to typeNode")
	t.Require().Equal(tStructName, tn2.name, "typeNode name is incorrect")
	t.Require().Equal(tn, tn2, "typeNode is not equal to original")
	t.Require().Equal("servicegraph.testStruct[name=\"testStruct\"]", fmt.Sprint(tn2), "typeNode string is incorrect")
}
