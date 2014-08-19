// This file was auto-generated by the veyron vdl tool.
// Source: poolheater.vdl

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

// PoolHeater allows clients to control when the pool is being heated.
// PoolHeater is the interface the client binds and uses.
// PoolHeater_ExcludingUniversal is the interface without internal framework-added methods
// to enable embedding without method collisions.  Not to be used directly by clients.
type PoolHeater_ExcludingUniversal interface {
	// Status retrieves the PoolHeater's status (i.e., active, idle)
	Status(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply string, err error)
	// Start informs the PoolHeater to heat the pool to the given temperature until the duration expires.
	Start(ctx _gen_context.T, temperature uint64, duration uint64, opts ..._gen_ipc.CallOpt) (err error)
	// Stop informs the PoolHeater to cease heating the pool.
	Stop(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (err error)
}
type PoolHeater interface {
	_gen_ipc.UniversalServiceMethods
	PoolHeater_ExcludingUniversal
}

// PoolHeaterService is the interface the server implements.
type PoolHeaterService interface {

	// Status retrieves the PoolHeater's status (i.e., active, idle)
	Status(context _gen_ipc.ServerContext) (reply string, err error)
	// Start informs the PoolHeater to heat the pool to the given temperature until the duration expires.
	Start(context _gen_ipc.ServerContext, temperature uint64, duration uint64) (err error)
	// Stop informs the PoolHeater to cease heating the pool.
	Stop(context _gen_ipc.ServerContext) (err error)
}

// BindPoolHeater returns the client stub implementing the PoolHeater
// interface.
//
// If no _gen_ipc.Client is specified, the default _gen_ipc.Client in the
// global Runtime is used.
func BindPoolHeater(name string, opts ..._gen_ipc.BindOpt) (PoolHeater, error) {
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
	stub := &clientStubPoolHeater{client: client, name: name}

	return stub, nil
}

// NewServerPoolHeater creates a new server stub.
//
// It takes a regular server implementing the PoolHeaterService
// interface, and returns a new server stub.
func NewServerPoolHeater(server PoolHeaterService) interface{} {
	return &ServerStubPoolHeater{
		service: server,
	}
}

// clientStubPoolHeater implements PoolHeater.
type clientStubPoolHeater struct {
	client _gen_ipc.Client
	name   string
}

func (__gen_c *clientStubPoolHeater) Status(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply string, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Status", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubPoolHeater) Start(ctx _gen_context.T, temperature uint64, duration uint64, opts ..._gen_ipc.CallOpt) (err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Start", []interface{}{temperature, duration}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubPoolHeater) Stop(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Stop", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubPoolHeater) UnresolveStep(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply []string, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "UnresolveStep", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubPoolHeater) Signature(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply _gen_ipc.ServiceSignature, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "Signature", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubPoolHeater) GetMethodTags(ctx _gen_context.T, method string, opts ..._gen_ipc.CallOpt) (reply []interface{}, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client.StartCall(ctx, __gen_c.name, "GetMethodTags", []interface{}{method}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

// ServerStubPoolHeater wraps a server that implements
// PoolHeaterService and provides an object that satisfies
// the requirements of veyron2/ipc.ReflectInvoker.
type ServerStubPoolHeater struct {
	service PoolHeaterService
}

func (__gen_s *ServerStubPoolHeater) GetMethodTags(call _gen_ipc.ServerCall, method string) ([]interface{}, error) {
	// TODO(bprosnitz) GetMethodTags() will be replaces with Signature().
	// Note: This exhibits some weird behavior like returning a nil error if the method isn't found.
	// This will change when it is replaced with Signature().
	switch method {
	case "Status":
		return []interface{}{}, nil
	case "Start":
		return []interface{}{}, nil
	case "Stop":
		return []interface{}{}, nil
	default:
		return nil, nil
	}
}

func (__gen_s *ServerStubPoolHeater) Signature(call _gen_ipc.ServerCall) (_gen_ipc.ServiceSignature, error) {
	result := _gen_ipc.ServiceSignature{Methods: make(map[string]_gen_ipc.MethodSignature)}
	result.Methods["Start"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "temperature", Type: 53},
			{Name: "duration", Type: 53},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 65},
		},
	}
	result.Methods["Status"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 3},
			{Name: "", Type: 65},
		},
	}
	result.Methods["Stop"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 65},
		},
	}

	result.TypeDefs = []_gen_vdlutil.Any{
		_gen_wiretype.NamedPrimitiveType{Type: 0x1, Name: "error", Tags: []string(nil)}}

	return result, nil
}

func (__gen_s *ServerStubPoolHeater) UnresolveStep(call _gen_ipc.ServerCall) (reply []string, err error) {
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

func (__gen_s *ServerStubPoolHeater) Status(call _gen_ipc.ServerCall) (reply string, err error) {
	reply, err = __gen_s.service.Status(call)
	return
}

func (__gen_s *ServerStubPoolHeater) Start(call _gen_ipc.ServerCall, temperature uint64, duration uint64) (err error) {
	err = __gen_s.service.Start(call, temperature, duration)
	return
}

func (__gen_s *ServerStubPoolHeater) Stop(call _gen_ipc.ServerCall) (err error) {
	err = __gen_s.service.Stop(call)
	return
}