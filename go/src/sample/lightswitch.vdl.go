// This file was auto-generated by the veyron vdl tool.
// Source: lightswitch.vdl

package sample

import (
	// The non-user imports are prefixed with "__" to prevent collisions.
	__veyron2 "veyron.io/veyron/veyron2"
	__context "veyron.io/veyron/veyron2/context"
	__ipc "veyron.io/veyron/veyron2/ipc"
	__vdlutil "veyron.io/veyron/veyron2/vdl/vdlutil"
	__wiretype "veyron.io/veyron/veyron2/wiretype"
)

// TODO(toddw): Remove this line once the new signature support is done.
// It corrects a bug where __wiretype is unused in VDL pacakges where only
// bootstrap types are used on interfaces.
const _ = __wiretype.TypeIDInvalid

// LightSwitchClientMethods is the client interface
// containing LightSwitch methods.
//
// LightSwitch allows clients to manipulate a virtual light switch.
type LightSwitchClientMethods interface {
	// Status indicates whether the light is on or off.
	Status(__context.T, ...__ipc.CallOpt) (string, error)
	// FlipSwitch sets the light to on or off, depending on the input.
	FlipSwitch(ctx __context.T, toOn bool, opts ...__ipc.CallOpt) error
}

// LightSwitchClientStub adds universal methods to LightSwitchClientMethods.
type LightSwitchClientStub interface {
	LightSwitchClientMethods
	__ipc.UniversalServiceMethods
}

// LightSwitchClient returns a client stub for LightSwitch.
func LightSwitchClient(name string, opts ...__ipc.BindOpt) LightSwitchClientStub {
	var client __ipc.Client
	for _, opt := range opts {
		if clientOpt, ok := opt.(__ipc.Client); ok {
			client = clientOpt
		}
	}
	return implLightSwitchClientStub{name, client}
}

type implLightSwitchClientStub struct {
	name   string
	client __ipc.Client
}

func (c implLightSwitchClientStub) c(ctx __context.T) __ipc.Client {
	if c.client != nil {
		return c.client
	}
	return __veyron2.RuntimeFromContext(ctx).Client()
}

func (c implLightSwitchClientStub) Status(ctx __context.T, opts ...__ipc.CallOpt) (o0 string, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Status", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&o0, &err); ierr != nil {
		err = ierr
	}
	return
}

func (c implLightSwitchClientStub) FlipSwitch(ctx __context.T, i0 bool, opts ...__ipc.CallOpt) (err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "FlipSwitch", []interface{}{i0}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (c implLightSwitchClientStub) Signature(ctx __context.T, opts ...__ipc.CallOpt) (o0 __ipc.ServiceSignature, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Signature", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&o0, &err); ierr != nil {
		err = ierr
	}
	return
}

// LightSwitchServerMethods is the interface a server writer
// implements for LightSwitch.
//
// LightSwitch allows clients to manipulate a virtual light switch.
type LightSwitchServerMethods interface {
	// Status indicates whether the light is on or off.
	Status(__ipc.ServerContext) (string, error)
	// FlipSwitch sets the light to on or off, depending on the input.
	FlipSwitch(ctx __ipc.ServerContext, toOn bool) error
}

// LightSwitchServerStubMethods is the server interface containing
// LightSwitch methods, as expected by ipc.Server.
// There is no difference between this interface and LightSwitchServerMethods
// since there are no streaming methods.
type LightSwitchServerStubMethods LightSwitchServerMethods

// LightSwitchServerStub adds universal methods to LightSwitchServerStubMethods.
type LightSwitchServerStub interface {
	LightSwitchServerStubMethods
	// Describe the LightSwitch interfaces.
	Describe__() []__ipc.InterfaceDesc
	// Signature will be replaced with Describe__.
	Signature(ctx __ipc.ServerContext) (__ipc.ServiceSignature, error)
}

// LightSwitchServer returns a server stub for LightSwitch.
// It converts an implementation of LightSwitchServerMethods into
// an object that may be used by ipc.Server.
func LightSwitchServer(impl LightSwitchServerMethods) LightSwitchServerStub {
	stub := implLightSwitchServerStub{
		impl: impl,
	}
	// Initialize GlobState; always check the stub itself first, to handle the
	// case where the user has the Glob method defined in their VDL source.
	if gs := __ipc.NewGlobState(stub); gs != nil {
		stub.gs = gs
	} else if gs := __ipc.NewGlobState(impl); gs != nil {
		stub.gs = gs
	}
	return stub
}

type implLightSwitchServerStub struct {
	impl LightSwitchServerMethods
	gs   *__ipc.GlobState
}

func (s implLightSwitchServerStub) Status(ctx __ipc.ServerContext) (string, error) {
	return s.impl.Status(ctx)
}

func (s implLightSwitchServerStub) FlipSwitch(ctx __ipc.ServerContext, i0 bool) error {
	return s.impl.FlipSwitch(ctx, i0)
}

func (s implLightSwitchServerStub) VGlob() *__ipc.GlobState {
	return s.gs
}

func (s implLightSwitchServerStub) Describe__() []__ipc.InterfaceDesc {
	return []__ipc.InterfaceDesc{LightSwitchDesc}
}

// LightSwitchDesc describes the LightSwitch interface.
var LightSwitchDesc __ipc.InterfaceDesc = descLightSwitch

// descLightSwitch hides the desc to keep godoc clean.
var descLightSwitch = __ipc.InterfaceDesc{
	Name:    "LightSwitch",
	PkgPath: "sample",
	Doc:     "// LightSwitch allows clients to manipulate a virtual light switch.",
	Methods: []__ipc.MethodDesc{
		{
			Name: "Status",
			Doc:  "// Status indicates whether the light is on or off.",
			OutArgs: []__ipc.ArgDesc{
				{"", ``}, // string
				{"", ``}, // error
			},
		},
		{
			Name: "FlipSwitch",
			Doc:  "// FlipSwitch sets the light to on or off, depending on the input.",
			InArgs: []__ipc.ArgDesc{
				{"toOn", ``}, // bool
			},
			OutArgs: []__ipc.ArgDesc{
				{"", ``}, // error
			},
		},
	},
}

func (s implLightSwitchServerStub) Signature(ctx __ipc.ServerContext) (__ipc.ServiceSignature, error) {
	// TODO(toddw): Replace with new Describe__ implementation.
	result := __ipc.ServiceSignature{Methods: make(map[string]__ipc.MethodSignature)}
	result.Methods["FlipSwitch"] = __ipc.MethodSignature{
		InArgs: []__ipc.MethodArgument{
			{Name: "toOn", Type: 2},
		},
		OutArgs: []__ipc.MethodArgument{
			{Name: "", Type: 65},
		},
	}
	result.Methods["Status"] = __ipc.MethodSignature{
		InArgs: []__ipc.MethodArgument{},
		OutArgs: []__ipc.MethodArgument{
			{Name: "", Type: 3},
			{Name: "", Type: 65},
		},
	}

	result.TypeDefs = []__vdlutil.Any{
		__wiretype.NamedPrimitiveType{Type: 0x1, Name: "error", Tags: []string(nil)}}

	return result, nil
}
