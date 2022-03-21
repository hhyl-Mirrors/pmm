// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package managementpb

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

// SecurityChecksClient is the client API for SecurityChecks service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SecurityChecksClient interface {
	// GetSecurityCheckResults returns Security Thread Tool's latest checks results.
	GetSecurityCheckResults(ctx context.Context, in *GetSecurityCheckResultsRequest, opts ...grpc.CallOption) (*GetSecurityCheckResultsResponse, error)
	// StartSecurityChecks executes Security Thread Tool checks and returns when all checks are executed.
	StartSecurityChecks(ctx context.Context, in *StartSecurityChecksRequest, opts ...grpc.CallOption) (*StartSecurityChecksResponse, error)
	// ListSecurityChecks returns a list of available Security Thread Tool checks.
	ListSecurityChecks(ctx context.Context, in *ListSecurityChecksRequest, opts ...grpc.CallOption) (*ListSecurityChecksResponse, error)
	// ChangeSecurityChecks enables/disables Security Thread Tool checks or changes their interval by names.
	ChangeSecurityChecks(ctx context.Context, in *ChangeSecurityChecksRequest, opts ...grpc.CallOption) (*ChangeSecurityChecksResponse, error)
}

type securityChecksClient struct {
	cc grpc.ClientConnInterface
}

func NewSecurityChecksClient(cc grpc.ClientConnInterface) SecurityChecksClient {
	return &securityChecksClient{cc}
}

func (c *securityChecksClient) GetSecurityCheckResults(ctx context.Context, in *GetSecurityCheckResultsRequest, opts ...grpc.CallOption) (*GetSecurityCheckResultsResponse, error) {
	out := new(GetSecurityCheckResultsResponse)
	err := c.cc.Invoke(ctx, "/management.SecurityChecks/GetSecurityCheckResults", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityChecksClient) StartSecurityChecks(ctx context.Context, in *StartSecurityChecksRequest, opts ...grpc.CallOption) (*StartSecurityChecksResponse, error) {
	out := new(StartSecurityChecksResponse)
	err := c.cc.Invoke(ctx, "/management.SecurityChecks/StartSecurityChecks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityChecksClient) ListSecurityChecks(ctx context.Context, in *ListSecurityChecksRequest, opts ...grpc.CallOption) (*ListSecurityChecksResponse, error) {
	out := new(ListSecurityChecksResponse)
	err := c.cc.Invoke(ctx, "/management.SecurityChecks/ListSecurityChecks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityChecksClient) ChangeSecurityChecks(ctx context.Context, in *ChangeSecurityChecksRequest, opts ...grpc.CallOption) (*ChangeSecurityChecksResponse, error) {
	out := new(ChangeSecurityChecksResponse)
	err := c.cc.Invoke(ctx, "/management.SecurityChecks/ChangeSecurityChecks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecurityChecksServer is the server API for SecurityChecks service.
// All implementations must embed UnimplementedSecurityChecksServer
// for forward compatibility
type SecurityChecksServer interface {
	// GetSecurityCheckResults returns Security Thread Tool's latest checks results.
	GetSecurityCheckResults(context.Context, *GetSecurityCheckResultsRequest) (*GetSecurityCheckResultsResponse, error)
	// StartSecurityChecks executes Security Thread Tool checks and returns when all checks are executed.
	StartSecurityChecks(context.Context, *StartSecurityChecksRequest) (*StartSecurityChecksResponse, error)
	// ListSecurityChecks returns a list of available Security Thread Tool checks.
	ListSecurityChecks(context.Context, *ListSecurityChecksRequest) (*ListSecurityChecksResponse, error)
	// ChangeSecurityChecks enables/disables Security Thread Tool checks or changes their interval by names.
	ChangeSecurityChecks(context.Context, *ChangeSecurityChecksRequest) (*ChangeSecurityChecksResponse, error)
	mustEmbedUnimplementedSecurityChecksServer()
}

// UnimplementedSecurityChecksServer must be embedded to have forward compatible implementations.
type UnimplementedSecurityChecksServer struct {
}

func (UnimplementedSecurityChecksServer) GetSecurityCheckResults(context.Context, *GetSecurityCheckResultsRequest) (*GetSecurityCheckResultsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSecurityCheckResults not implemented")
}
func (UnimplementedSecurityChecksServer) StartSecurityChecks(context.Context, *StartSecurityChecksRequest) (*StartSecurityChecksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartSecurityChecks not implemented")
}
func (UnimplementedSecurityChecksServer) ListSecurityChecks(context.Context, *ListSecurityChecksRequest) (*ListSecurityChecksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSecurityChecks not implemented")
}
func (UnimplementedSecurityChecksServer) ChangeSecurityChecks(context.Context, *ChangeSecurityChecksRequest) (*ChangeSecurityChecksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeSecurityChecks not implemented")
}
func (UnimplementedSecurityChecksServer) mustEmbedUnimplementedSecurityChecksServer() {}

// UnsafeSecurityChecksServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SecurityChecksServer will
// result in compilation errors.
type UnsafeSecurityChecksServer interface {
	mustEmbedUnimplementedSecurityChecksServer()
}

func RegisterSecurityChecksServer(s grpc.ServiceRegistrar, srv SecurityChecksServer) {
	s.RegisterService(&SecurityChecks_ServiceDesc, srv)
}

func _SecurityChecks_GetSecurityCheckResults_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSecurityCheckResultsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityChecksServer).GetSecurityCheckResults(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.SecurityChecks/GetSecurityCheckResults",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityChecksServer).GetSecurityCheckResults(ctx, req.(*GetSecurityCheckResultsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityChecks_StartSecurityChecks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartSecurityChecksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityChecksServer).StartSecurityChecks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.SecurityChecks/StartSecurityChecks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityChecksServer).StartSecurityChecks(ctx, req.(*StartSecurityChecksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityChecks_ListSecurityChecks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSecurityChecksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityChecksServer).ListSecurityChecks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.SecurityChecks/ListSecurityChecks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityChecksServer).ListSecurityChecks(ctx, req.(*ListSecurityChecksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityChecks_ChangeSecurityChecks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeSecurityChecksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityChecksServer).ChangeSecurityChecks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.SecurityChecks/ChangeSecurityChecks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityChecksServer).ChangeSecurityChecks(ctx, req.(*ChangeSecurityChecksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SecurityChecks_ServiceDesc is the grpc.ServiceDesc for SecurityChecks service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SecurityChecks_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "management.SecurityChecks",
	HandlerType: (*SecurityChecksServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSecurityCheckResults",
			Handler:    _SecurityChecks_GetSecurityCheckResults_Handler,
		},
		{
			MethodName: "StartSecurityChecks",
			Handler:    _SecurityChecks_StartSecurityChecks_Handler,
		},
		{
			MethodName: "ListSecurityChecks",
			Handler:    _SecurityChecks_ListSecurityChecks_Handler,
		},
		{
			MethodName: "ChangeSecurityChecks",
			Handler:    _SecurityChecks_ChangeSecurityChecks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "managementpb/checks.proto",
}
