// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package inventorypb

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

// ServicesClient is the client API for Services service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServicesClient interface {
	// ListServices returns a list of Services filtered by type.
	ListServices(ctx context.Context, in *ListServicesRequest, opts ...grpc.CallOption) (*ListServicesResponse, error)
	// GetService returns a single Service by ID.
	GetService(ctx context.Context, in *GetServiceRequest, opts ...grpc.CallOption) (*GetServiceResponse, error)
	// AddMySQLService adds MySQL Service.
	AddMySQLService(ctx context.Context, in *AddMySQLServiceRequest, opts ...grpc.CallOption) (*AddMySQLServiceResponse, error)
	// AddMongoDBService adds MongoDB Service.
	AddMongoDBService(ctx context.Context, in *AddMongoDBServiceRequest, opts ...grpc.CallOption) (*AddMongoDBServiceResponse, error)
	// AddPostgreSQLService adds PostgreSQL Service.
	AddPostgreSQLService(ctx context.Context, in *AddPostgreSQLServiceRequest, opts ...grpc.CallOption) (*AddPostgreSQLServiceResponse, error)
	// AddProxySQLService adds ProxySQL Service.
	AddProxySQLService(ctx context.Context, in *AddProxySQLServiceRequest, opts ...grpc.CallOption) (*AddProxySQLServiceResponse, error)
	// AddHAProxyService adds HAProxy Service.
	AddHAProxyService(ctx context.Context, in *AddHAProxyServiceRequest, opts ...grpc.CallOption) (*AddHAProxyServiceResponse, error)
	// AddExternalService adds External Service.
	AddExternalService(ctx context.Context, in *AddExternalServiceRequest, opts ...grpc.CallOption) (*AddExternalServiceResponse, error)
	// RemoveService removes Service.
	RemoveService(ctx context.Context, in *RemoveServiceRequest, opts ...grpc.CallOption) (*RemoveServiceResponse, error)
}

type servicesClient struct {
	cc grpc.ClientConnInterface
}

func NewServicesClient(cc grpc.ClientConnInterface) ServicesClient {
	return &servicesClient{cc}
}

func (c *servicesClient) ListServices(ctx context.Context, in *ListServicesRequest, opts ...grpc.CallOption) (*ListServicesResponse, error) {
	out := new(ListServicesResponse)
	err := c.cc.Invoke(ctx, "/inventory.Services/ListServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesClient) GetService(ctx context.Context, in *GetServiceRequest, opts ...grpc.CallOption) (*GetServiceResponse, error) {
	out := new(GetServiceResponse)
	err := c.cc.Invoke(ctx, "/inventory.Services/GetService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesClient) AddMySQLService(ctx context.Context, in *AddMySQLServiceRequest, opts ...grpc.CallOption) (*AddMySQLServiceResponse, error) {
	out := new(AddMySQLServiceResponse)
	err := c.cc.Invoke(ctx, "/inventory.Services/AddMySQLService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesClient) AddMongoDBService(ctx context.Context, in *AddMongoDBServiceRequest, opts ...grpc.CallOption) (*AddMongoDBServiceResponse, error) {
	out := new(AddMongoDBServiceResponse)
	err := c.cc.Invoke(ctx, "/inventory.Services/AddMongoDBService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesClient) AddPostgreSQLService(ctx context.Context, in *AddPostgreSQLServiceRequest, opts ...grpc.CallOption) (*AddPostgreSQLServiceResponse, error) {
	out := new(AddPostgreSQLServiceResponse)
	err := c.cc.Invoke(ctx, "/inventory.Services/AddPostgreSQLService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesClient) AddProxySQLService(ctx context.Context, in *AddProxySQLServiceRequest, opts ...grpc.CallOption) (*AddProxySQLServiceResponse, error) {
	out := new(AddProxySQLServiceResponse)
	err := c.cc.Invoke(ctx, "/inventory.Services/AddProxySQLService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesClient) AddHAProxyService(ctx context.Context, in *AddHAProxyServiceRequest, opts ...grpc.CallOption) (*AddHAProxyServiceResponse, error) {
	out := new(AddHAProxyServiceResponse)
	err := c.cc.Invoke(ctx, "/inventory.Services/AddHAProxyService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesClient) AddExternalService(ctx context.Context, in *AddExternalServiceRequest, opts ...grpc.CallOption) (*AddExternalServiceResponse, error) {
	out := new(AddExternalServiceResponse)
	err := c.cc.Invoke(ctx, "/inventory.Services/AddExternalService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesClient) RemoveService(ctx context.Context, in *RemoveServiceRequest, opts ...grpc.CallOption) (*RemoveServiceResponse, error) {
	out := new(RemoveServiceResponse)
	err := c.cc.Invoke(ctx, "/inventory.Services/RemoveService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServicesServer is the server API for Services service.
// All implementations must embed UnimplementedServicesServer
// for forward compatibility
type ServicesServer interface {
	// ListServices returns a list of Services filtered by type.
	ListServices(context.Context, *ListServicesRequest) (*ListServicesResponse, error)
	// GetService returns a single Service by ID.
	GetService(context.Context, *GetServiceRequest) (*GetServiceResponse, error)
	// AddMySQLService adds MySQL Service.
	AddMySQLService(context.Context, *AddMySQLServiceRequest) (*AddMySQLServiceResponse, error)
	// AddMongoDBService adds MongoDB Service.
	AddMongoDBService(context.Context, *AddMongoDBServiceRequest) (*AddMongoDBServiceResponse, error)
	// AddPostgreSQLService adds PostgreSQL Service.
	AddPostgreSQLService(context.Context, *AddPostgreSQLServiceRequest) (*AddPostgreSQLServiceResponse, error)
	// AddProxySQLService adds ProxySQL Service.
	AddProxySQLService(context.Context, *AddProxySQLServiceRequest) (*AddProxySQLServiceResponse, error)
	// AddHAProxyService adds HAProxy Service.
	AddHAProxyService(context.Context, *AddHAProxyServiceRequest) (*AddHAProxyServiceResponse, error)
	// AddExternalService adds External Service.
	AddExternalService(context.Context, *AddExternalServiceRequest) (*AddExternalServiceResponse, error)
	// RemoveService removes Service.
	RemoveService(context.Context, *RemoveServiceRequest) (*RemoveServiceResponse, error)
	mustEmbedUnimplementedServicesServer()
}

// UnimplementedServicesServer must be embedded to have forward compatible implementations.
type UnimplementedServicesServer struct {
}

func (UnimplementedServicesServer) ListServices(context.Context, *ListServicesRequest) (*ListServicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListServices not implemented")
}
func (UnimplementedServicesServer) GetService(context.Context, *GetServiceRequest) (*GetServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetService not implemented")
}
func (UnimplementedServicesServer) AddMySQLService(context.Context, *AddMySQLServiceRequest) (*AddMySQLServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMySQLService not implemented")
}
func (UnimplementedServicesServer) AddMongoDBService(context.Context, *AddMongoDBServiceRequest) (*AddMongoDBServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMongoDBService not implemented")
}
func (UnimplementedServicesServer) AddPostgreSQLService(context.Context, *AddPostgreSQLServiceRequest) (*AddPostgreSQLServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPostgreSQLService not implemented")
}
func (UnimplementedServicesServer) AddProxySQLService(context.Context, *AddProxySQLServiceRequest) (*AddProxySQLServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProxySQLService not implemented")
}
func (UnimplementedServicesServer) AddHAProxyService(context.Context, *AddHAProxyServiceRequest) (*AddHAProxyServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddHAProxyService not implemented")
}
func (UnimplementedServicesServer) AddExternalService(context.Context, *AddExternalServiceRequest) (*AddExternalServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddExternalService not implemented")
}
func (UnimplementedServicesServer) RemoveService(context.Context, *RemoveServiceRequest) (*RemoveServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveService not implemented")
}
func (UnimplementedServicesServer) mustEmbedUnimplementedServicesServer() {}

// UnsafeServicesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServicesServer will
// result in compilation errors.
type UnsafeServicesServer interface {
	mustEmbedUnimplementedServicesServer()
}

func RegisterServicesServer(s grpc.ServiceRegistrar, srv ServicesServer) {
	s.RegisterService(&Services_ServiceDesc, srv)
}

func _Services_ListServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListServicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServer).ListServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inventory.Services/ListServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServer).ListServices(ctx, req.(*ListServicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Services_GetService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServer).GetService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inventory.Services/GetService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServer).GetService(ctx, req.(*GetServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Services_AddMySQLService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMySQLServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServer).AddMySQLService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inventory.Services/AddMySQLService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServer).AddMySQLService(ctx, req.(*AddMySQLServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Services_AddMongoDBService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMongoDBServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServer).AddMongoDBService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inventory.Services/AddMongoDBService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServer).AddMongoDBService(ctx, req.(*AddMongoDBServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Services_AddPostgreSQLService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPostgreSQLServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServer).AddPostgreSQLService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inventory.Services/AddPostgreSQLService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServer).AddPostgreSQLService(ctx, req.(*AddPostgreSQLServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Services_AddProxySQLService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddProxySQLServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServer).AddProxySQLService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inventory.Services/AddProxySQLService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServer).AddProxySQLService(ctx, req.(*AddProxySQLServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Services_AddHAProxyService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddHAProxyServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServer).AddHAProxyService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inventory.Services/AddHAProxyService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServer).AddHAProxyService(ctx, req.(*AddHAProxyServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Services_AddExternalService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddExternalServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServer).AddExternalService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inventory.Services/AddExternalService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServer).AddExternalService(ctx, req.(*AddExternalServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Services_RemoveService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServer).RemoveService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inventory.Services/RemoveService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServer).RemoveService(ctx, req.(*RemoveServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Services_ServiceDesc is the grpc.ServiceDesc for Services service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Services_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inventory.Services",
	HandlerType: (*ServicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListServices",
			Handler:    _Services_ListServices_Handler,
		},
		{
			MethodName: "GetService",
			Handler:    _Services_GetService_Handler,
		},
		{
			MethodName: "AddMySQLService",
			Handler:    _Services_AddMySQLService_Handler,
		},
		{
			MethodName: "AddMongoDBService",
			Handler:    _Services_AddMongoDBService_Handler,
		},
		{
			MethodName: "AddPostgreSQLService",
			Handler:    _Services_AddPostgreSQLService_Handler,
		},
		{
			MethodName: "AddProxySQLService",
			Handler:    _Services_AddProxySQLService_Handler,
		},
		{
			MethodName: "AddHAProxyService",
			Handler:    _Services_AddHAProxyService_Handler,
		},
		{
			MethodName: "AddExternalService",
			Handler:    _Services_AddExternalService_Handler,
		},
		{
			MethodName: "RemoveService",
			Handler:    _Services_RemoveService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "inventorypb/services.proto",
}
