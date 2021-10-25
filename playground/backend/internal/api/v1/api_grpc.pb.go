//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package playground

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PlaygroundServiceClient is the client API for PlaygroundService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PlaygroundServiceClient interface {
	// Submit the job for an execution and get the pipeline uuid.
	RunCode(ctx context.Context, in *RunCodeRequest, opts ...grpc.CallOption) (*RunCodeResponse, error)
	// Get the status of pipeline execution.
	CheckStatus(ctx context.Context, in *CheckStatusRequest, opts ...grpc.CallOption) (*CheckStatusResponse, error)
	// Get the result of pipeline execution.
	GetRunOutput(ctx context.Context, in *GetRunOutputRequest, opts ...grpc.CallOption) (*GetRunOutputResponse, error)
	// Get the result of pipeline compilation.
	GetCompileOutput(ctx context.Context, in *GetCompileOutputRequest, opts ...grpc.CallOption) (*GetCompileOutputResponse, error)
}

type playgroundServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPlaygroundServiceClient(cc grpc.ClientConnInterface) PlaygroundServiceClient {
	return &playgroundServiceClient{cc}
}

func (c *playgroundServiceClient) RunCode(ctx context.Context, in *RunCodeRequest, opts ...grpc.CallOption) (*RunCodeResponse, error) {
	out := new(RunCodeResponse)
	err := c.cc.Invoke(ctx, "/api.v1.PlaygroundService/RunCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playgroundServiceClient) CheckStatus(ctx context.Context, in *CheckStatusRequest, opts ...grpc.CallOption) (*CheckStatusResponse, error) {
	out := new(CheckStatusResponse)
	err := c.cc.Invoke(ctx, "/api.v1.PlaygroundService/CheckStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playgroundServiceClient) GetRunOutput(ctx context.Context, in *GetRunOutputRequest, opts ...grpc.CallOption) (*GetRunOutputResponse, error) {
	out := new(GetRunOutputResponse)
	err := c.cc.Invoke(ctx, "/api.v1.PlaygroundService/GetRunOutput", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playgroundServiceClient) GetCompileOutput(ctx context.Context, in *GetCompileOutputRequest, opts ...grpc.CallOption) (*GetCompileOutputResponse, error) {
	out := new(GetCompileOutputResponse)
	err := c.cc.Invoke(ctx, "/api.v1.PlaygroundService/GetCompileOutput", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlaygroundServiceServer is the server API for PlaygroundService service.
// All implementations must embed UnimplementedPlaygroundServiceServer
// for forward compatibility
type PlaygroundServiceServer interface {
	// Submit the job for an execution and get the pipeline uuid.
	RunCode(context.Context, *RunCodeRequest) (*RunCodeResponse, error)
	// Get the status of pipeline execution.
	CheckStatus(context.Context, *CheckStatusRequest) (*CheckStatusResponse, error)
	// Get the result of pipeline execution.
	GetRunOutput(context.Context, *GetRunOutputRequest) (*GetRunOutputResponse, error)
	// Get the result of pipeline compilation.
	GetCompileOutput(context.Context, *GetCompileOutputRequest) (*GetCompileOutputResponse, error)
	mustEmbedUnimplementedPlaygroundServiceServer()
}

// UnimplementedPlaygroundServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPlaygroundServiceServer struct {
}

func (UnimplementedPlaygroundServiceServer) RunCode(context.Context, *RunCodeRequest) (*RunCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunCode not implemented")
}
func (UnimplementedPlaygroundServiceServer) CheckStatus(context.Context, *CheckStatusRequest) (*CheckStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckStatus not implemented")
}
func (UnimplementedPlaygroundServiceServer) GetRunOutput(context.Context, *GetRunOutputRequest) (*GetRunOutputResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRunOutput not implemented")
}
func (UnimplementedPlaygroundServiceServer) GetCompileOutput(context.Context, *GetCompileOutputRequest) (*GetCompileOutputResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompileOutput not implemented")
}
func (UnimplementedPlaygroundServiceServer) mustEmbedUnimplementedPlaygroundServiceServer() {}

// UnsafePlaygroundServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PlaygroundServiceServer will
// result in compilation errors.
type UnsafePlaygroundServiceServer interface {
	mustEmbedUnimplementedPlaygroundServiceServer()
}

func RegisterPlaygroundServiceServer(s grpc.ServiceRegistrar, srv PlaygroundServiceServer) {
	s.RegisterService(&PlaygroundService_ServiceDesc, srv)
}

func _PlaygroundService_RunCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaygroundServiceServer).RunCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.PlaygroundService/RunCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaygroundServiceServer).RunCode(ctx, req.(*RunCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlaygroundService_CheckStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaygroundServiceServer).CheckStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.PlaygroundService/CheckStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaygroundServiceServer).CheckStatus(ctx, req.(*CheckStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlaygroundService_GetRunOutput_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRunOutputRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaygroundServiceServer).GetRunOutput(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.PlaygroundService/GetRunOutput",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaygroundServiceServer).GetRunOutput(ctx, req.(*GetRunOutputRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlaygroundService_GetCompileOutput_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompileOutputRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaygroundServiceServer).GetCompileOutput(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.PlaygroundService/GetCompileOutput",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaygroundServiceServer).GetCompileOutput(ctx, req.(*GetCompileOutputRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PlaygroundService_ServiceDesc is the grpc.ServiceDesc for PlaygroundService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PlaygroundService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.PlaygroundService",
	HandlerType: (*PlaygroundServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RunCode",
			Handler:    _PlaygroundService_RunCode_Handler,
		},
		{
			MethodName: "CheckStatus",
			Handler:    _PlaygroundService_CheckStatus_Handler,
		},
		{
			MethodName: "GetRunOutput",
			Handler:    _PlaygroundService_GetRunOutput_Handler,
		},
		{
			MethodName: "GetCompileOutput",
			Handler:    _PlaygroundService_GetCompileOutput_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
