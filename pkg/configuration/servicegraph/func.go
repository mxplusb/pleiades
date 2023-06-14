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
	"net/url"
	"reflect"
	"runtime"
	"strings"
)

// funcType contains runtime information about a function.
type funcType struct {
	// Name of the function.
	name string
	// Name of the package in which this function is defined.
	pakage string

	// Path to the file in which this function is defined.
	file string

	// Line number in the file at which this function is defined.
	line int
}

// String returns a string representation of the function.
func (f *funcType) String() string {
	return fmt.Sprint(f)
}

// Format implements fmt.Formatter for Func, printing a single-line
// representation for %v and a multi-line one for %+v.
func (f *funcType) Format(w fmt.State, c rune) {
	if w.Flag('+') && c == 'v' {
		// "path/to/package".MyFunction
		// 	path/to/file.go:42
		fmt.Fprintf(w, "%q.%v", f.pakage, f.name)
		fmt.Fprintf(w, "\n\t%v:%v", f.file, f.line)
	} else {
		// "path/to/package".MyFunction (path/to/file.go:42)
		fmt.Fprintf(w, "%q.%v (%v:%v)", f.pakage, f.name, f.file, f.line)
	}
}

// InspectFunc inspects and returns runtime information about the given
// function.
func inspectFunc(function interface{}) *funcType {
	fptr := reflect.ValueOf(function).Pointer()
	return inspectFuncPC(fptr)
}

// inspectFuncPC inspects and returns runtime information about the function
// at the given program counter address.
func inspectFuncPC(pc uintptr) *funcType {
	f := runtime.FuncForPC(pc)
	if f == nil {
		return nil
	}
	pkgName, funcName := splitFuncName(f.Name())
	fileName, lineNum := f.FileLine(pc)
	return &funcType{
		name:    funcName,
		pakage: pkgName,
		file:    fileName,
		line:    lineNum,
	}
}

const _vendor = "/vendor/"

func splitFuncName(function string) (pname string, fname string) {
	if len(function) == 0 {
		return
	}

	// We have something like "path.to/my/pkg.MyFunction". If the function is
	// a closure, it is something like, "path.to/my/pkg.MyFunction.func1".

	idx := 0

	// Everything up to the first "." after the last "/" is the package name.
	// Everything after the "." is the full function name.
	if i := strings.LastIndex(function, "/"); i >= 0 {
		idx = i
	}
	if i := strings.Index(function[idx:], "."); i >= 0 {
		idx += i
	}
	pname, fname = function[:idx], function[idx+1:]

	// The package may be vendored.
	if i := strings.Index(pname, _vendor); i > 0 {
		pname = pname[i+len(_vendor):]
	}

	// Package names are URL-encoded to avoid ambiguity in the case where the
	// package name contains ".git". Otherwise, "foo/bar.git.MyFunction" would
	// mean that "git" is the top-level function and "MyFunction" is embedded
	// inside it.
	if unescaped, err := url.QueryUnescape(pname); err == nil {
		pname = unescaped
	}

	return
}
