// This file was auto-generated by the veyron vdl tool.
// Source: window.vdl

package sample

import (
	// The non-user imports are prefixed with "_gen_" to prevent collisions.
	_gen_context "veyron2/context"
	_gen_ipc "veyron2/ipc"
	_gen_naming "veyron2/naming"
	_gen_rt "veyron2/rt"
	_gen_vdlutil "veyron2/vdl/vdlutil"
	_gen_wiretype "veyron2/wiretype"
)

// TODO(bprosnitz) Remove this line once signatures are updated to use typevals.
// It corrects a bug where _gen_wiretype is unused in VDL pacakges where only bootstrap types are used on interfaces.
const _ = _gen_wiretype.TypeIDInvalid

// Window is the interface the client binds and uses.
// Window_ExcludingUniversal is the interface without internal framework-added methods
// to enable embedding without method collisions.  Not to be used directly by clients.
type Window_ExcludingUniversal interface {
	Open(ctx _gen_context.T, degree int16, opts ..._gen_ipc.CallOpt) (err error)
	OpenFully(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (err error)
	Close(ctx _gen_context.T, degree int16, opts ..._gen_ipc.CallOpt) (err error)
	CloseFully(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (err error)
}
type Window interface {
	_gen_ipc.UniversalServiceMethods
	Window_ExcludingUniversal
}

// WindowService is the interface the server implements.
type WindowService interface {
	Open(context _gen_ipc.ServerContext, degree int16) (err error)
	OpenFully(context _gen_ipc.ServerContext) (err error)
	Close(context _gen_ipc.ServerContext, degree int16) (err error)
	CloseFully(context _gen_ipc.ServerContext) (err error)
}

// BindWindow returns the client stub implementing the Window
// interface.
//
// If no _gen_ipc.Client is specified, the default _gen_ipc.Client in the
// global Runtime is used.
func BindWindow(name string, opts ..._gen_ipc.BindOpt) (Window, error) {
	var client _gen_ipc.Client
	switch len(opts) {
	case 0:
		client = _gen_rt.R().Client()
	case 1:
		switch o := opts[0].(type) {
		case _gen_ipc.Client:
			client = o
		default:
			return nil, _gen_vdlutil.ErrUnrecognizedOption
		}
	default:
		return nil, _gen_vdlutil.ErrTooManyOptionsToBind
	}
	stub := &clientStubWindow{client: client, name: name}

	return stub, nil
}

// NewServerWindow creates a new server stub.
//
// It takes a regular server implementing the WindowService
// interface, and returns a new server stub.
func NewServerWindow(server WindowService) interface{} {
	return &ServerStubWindow{
		service: server,
	}
}

// clientStubWindow implements Window.
type clientStubWindow struct {
	client _gen_ipc.Client
	name   string
}

func (__gen_c *clientStubWindow) Open(ctx _gen_context.T, degree int16, opts ..._gen_ipc.CallOpt) (err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Open", []interface{}{degree}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubWindow) OpenFully(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "OpenFully", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubWindow) Close(ctx _gen_context.T, degree int16, opts ..._gen_ipc.CallOpt) (err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Close", []interface{}{degree}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubWindow) CloseFully(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "CloseFully", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubWindow) UnresolveStep(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply []string, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "UnresolveStep", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubWindow) Signature(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply _gen_ipc.ServiceSignature, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Signature", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubWindow) GetMethodTags(ctx _gen_context.T, method string, opts ..._gen_ipc.CallOpt) (reply []interface{}, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "GetMethodTags", []interface{}{method}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

// ServerStubWindow wraps a server that implements
// WindowService and provides an object that satisfies
// the requirements of veyron2/ipc.ReflectInvoker.
type ServerStubWindow struct {
	service WindowService
}

func (__gen_s *ServerStubWindow) GetMethodTags(call _gen_ipc.ServerCall, method string) ([]interface{}, error) {
	// TODO(bprosnitz) GetMethodTags() will be replaces with Signature().
	// Note: This exhibits some weird behavior like returning a nil error if the method isn't found.
	// This will change when it is replaced with Signature().
	switch method {
	case "Open":
		return []interface{}{}, nil
	case "OpenFully":
		return []interface{}{}, nil
	case "Close":
		return []interface{}{}, nil
	case "CloseFully":
		return []interface{}{}, nil
	default:
		return nil, nil
	}
}

func (__gen_s *ServerStubWindow) Signature(call _gen_ipc.ServerCall) (_gen_ipc.ServiceSignature, error) {
	result := _gen_ipc.ServiceSignature{Methods: make(map[string]_gen_ipc.MethodSignature)}
	result.Methods["Close"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "degree", Type: 35},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 65},
		},
	}
	result.Methods["CloseFully"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 65},
		},
	}
	result.Methods["Open"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "degree", Type: 35},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 65},
		},
	}
	result.Methods["OpenFully"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 65},
		},
	}

	result.TypeDefs = []_gen_vdlutil.Any{
		_gen_wiretype.NamedPrimitiveType{Type: 0x1, Name: "error", Tags: []string(nil)}}

	return result, nil
}

func (__gen_s *ServerStubWindow) UnresolveStep(call _gen_ipc.ServerCall) (reply []string, err error) {
	if unresolver, ok := __gen_s.service.(_gen_ipc.Unresolver); ok {
		return unresolver.UnresolveStep(call)
	}
	if call.Server() == nil {
		return
	}
	var published []string
	if published, err = call.Server().Published(); err != nil || published == nil {
		return
	}
	reply = make([]string, len(published))
	for i, p := range published {
		reply[i] = _gen_naming.Join(p, call.Name())
	}
	return
}

func (__gen_s *ServerStubWindow) Open(call _gen_ipc.ServerCall, degree int16) (err error) {
	err = __gen_s.service.Open(call, degree)
	return
}

func (__gen_s *ServerStubWindow) OpenFully(call _gen_ipc.ServerCall) (err error) {
	err = __gen_s.service.OpenFully(call)
	return
}

func (__gen_s *ServerStubWindow) Close(call _gen_ipc.ServerCall, degree int16) (err error) {
	err = __gen_s.service.Close(call, degree)
	return
}

func (__gen_s *ServerStubWindow) CloseFully(call _gen_ipc.ServerCall) (err error) {
	err = __gen_s.service.CloseFully(call)
	return
}