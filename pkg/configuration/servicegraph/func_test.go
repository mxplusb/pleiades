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
)

func Test_splitFuncName(t *testing.T) {
	type args struct {
		function string
	}
	tests := []struct {
		name      string
		args      args
		wantPname string
		wantFname string
	}{
		{
			name: "Test splitFuncName",
			args: args{
				function: "github.com/mxplusb/pleiades/pkg/configuration/servicegraph.(*ServiceManager).AddOption",
			},
			wantPname: "github.com/mxplusb/pleiades/pkg/configuration/servicegraph",
			wantFname: "(*ServiceManager).AddOption",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPname, gotFname := splitFuncName(tt.args.function)
			if gotPname != tt.wantPname {
				t.Errorf("splitFuncName() gotPname = %v, want %v", gotPname, tt.wantPname)
			}
			if gotFname != tt.wantFname {
				t.Errorf("splitFuncName() gotFname = %v, want %v", gotFname, tt.wantFname)
			}
		})
	}
}
