// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: alarm.vdl

package sample

import (
	// VDL system imports
	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/rpc"
)

// AlarmClientMethods is the client interface
// containing Alarm methods.
//
// Alarm allows clients to manipulate an alarm and query its status.
type AlarmClientMethods interface {
	// Status returns the current status of the Alarm (i.e., armed, unarmed, panicking).
	Status(*context.T, ...rpc.CallOpt) (string, error)
	// Arm sets the Alarm to the armed state.
	Arm(*context.T, ...rpc.CallOpt) error
	// DelayArm sets the Alarm to the armed state after the given delay in seconds.
	DelayArm(ctx *context.T, seconds float32, opts ...rpc.CallOpt) error
	// Unarm sets the Alarm to the unarmed state.
	Unarm(*context.T, ...rpc.CallOpt) error
	// Panic sets the Alarm to the panicking state.
	Panic(*context.T, ...rpc.CallOpt) error
}

// AlarmClientStub adds universal methods to AlarmClientMethods.
type AlarmClientStub interface {
	AlarmClientMethods
	rpc.UniversalServiceMethods
}

// AlarmClient returns a client stub for Alarm.
func AlarmClient(name string) AlarmClientStub {
	return implAlarmClientStub{name}
}

type implAlarmClientStub struct {
	name string
}

func (c implAlarmClientStub) Status(ctx *context.T, opts ...rpc.CallOpt) (o0 string, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "Status", nil, []interface{}{&o0}, opts...)
	return
}

func (c implAlarmClientStub) Arm(ctx *context.T, opts ...rpc.CallOpt) (err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "Arm", nil, nil, opts...)
	return
}

func (c implAlarmClientStub) DelayArm(ctx *context.T, i0 float32, opts ...rpc.CallOpt) (err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "DelayArm", []interface{}{i0}, nil, opts...)
	return
}

func (c implAlarmClientStub) Unarm(ctx *context.T, opts ...rpc.CallOpt) (err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "Unarm", nil, nil, opts...)
	return
}

func (c implAlarmClientStub) Panic(ctx *context.T, opts ...rpc.CallOpt) (err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "Panic", nil, nil, opts...)
	return
}

// AlarmServerMethods is the interface a server writer
// implements for Alarm.
//
// Alarm allows clients to manipulate an alarm and query its status.
type AlarmServerMethods interface {
	// Status returns the current status of the Alarm (i.e., armed, unarmed, panicking).
	Status(rpc.ServerCall) (string, error)
	// Arm sets the Alarm to the armed state.
	Arm(rpc.ServerCall) error
	// DelayArm sets the Alarm to the armed state after the given delay in seconds.
	DelayArm(call rpc.ServerCall, seconds float32) error
	// Unarm sets the Alarm to the unarmed state.
	Unarm(rpc.ServerCall) error
	// Panic sets the Alarm to the panicking state.
	Panic(rpc.ServerCall) error
}

// AlarmServerStubMethods is the server interface containing
// Alarm methods, as expected by rpc.Server.
// There is no difference between this interface and AlarmServerMethods
// since there are no streaming methods.
type AlarmServerStubMethods AlarmServerMethods

// AlarmServerStub adds universal methods to AlarmServerStubMethods.
type AlarmServerStub interface {
	AlarmServerStubMethods
	// Describe the Alarm interfaces.
	Describe__() []rpc.InterfaceDesc
}

// AlarmServer returns a server stub for Alarm.
// It converts an implementation of AlarmServerMethods into
// an object that may be used by rpc.Server.
func AlarmServer(impl AlarmServerMethods) AlarmServerStub {
	stub := implAlarmServerStub{
		impl: impl,
	}
	// Initialize GlobState; always check the stub itself first, to handle the
	// case where the user has the Glob method defined in their VDL source.
	if gs := rpc.NewGlobState(stub); gs != nil {
		stub.gs = gs
	} else if gs := rpc.NewGlobState(impl); gs != nil {
		stub.gs = gs
	}
	return stub
}

type implAlarmServerStub struct {
	impl AlarmServerMethods
	gs   *rpc.GlobState
}

func (s implAlarmServerStub) Status(call rpc.ServerCall) (string, error) {
	return s.impl.Status(call)
}

func (s implAlarmServerStub) Arm(call rpc.ServerCall) error {
	return s.impl.Arm(call)
}

func (s implAlarmServerStub) DelayArm(call rpc.ServerCall, i0 float32) error {
	return s.impl.DelayArm(call, i0)
}

func (s implAlarmServerStub) Unarm(call rpc.ServerCall) error {
	return s.impl.Unarm(call)
}

func (s implAlarmServerStub) Panic(call rpc.ServerCall) error {
	return s.impl.Panic(call)
}

func (s implAlarmServerStub) Globber() *rpc.GlobState {
	return s.gs
}

func (s implAlarmServerStub) Describe__() []rpc.InterfaceDesc {
	return []rpc.InterfaceDesc{AlarmDesc}
}

// AlarmDesc describes the Alarm interface.
var AlarmDesc rpc.InterfaceDesc = descAlarm

// descAlarm hides the desc to keep godoc clean.
var descAlarm = rpc.InterfaceDesc{
	Name:    "Alarm",
	PkgPath: "v.io/x/browser/sample",
	Doc:     "// Alarm allows clients to manipulate an alarm and query its status.",
	Methods: []rpc.MethodDesc{
		{
			Name: "Status",
			Doc:  "// Status returns the current status of the Alarm (i.e., armed, unarmed, panicking).",
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // string
			},
		},
		{
			Name: "Arm",
			Doc:  "// Arm sets the Alarm to the armed state.",
		},
		{
			Name: "DelayArm",
			Doc:  "// DelayArm sets the Alarm to the armed state after the given delay in seconds.",
			InArgs: []rpc.ArgDesc{
				{"seconds", ``}, // float32
			},
		},
		{
			Name: "Unarm",
			Doc:  "// Unarm sets the Alarm to the unarmed state.",
		},
		{
			Name: "Panic",
			Doc:  "// Panic sets the Alarm to the panicking state.",
		},
	},
}
