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

// MongoDBClient is the client API for MongoDB service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MongoDBClient interface {
	// AddMongoDB adds MongoDB Service and starts several Agents.
	// It automatically adds a service to inventory, which is running on provided "node_id",
	// then adds "mongodb_exporter", and "qan_mongodb_profiler" agents
	// with provided "pmm_agent_id" and other parameters.
	AddMongoDB(ctx context.Context, in *AddMongoDBRequest, opts ...grpc.CallOption) (*AddMongoDBResponse, error)
}

type mongoDBClient struct {
	cc grpc.ClientConnInterface
}

func NewMongoDBClient(cc grpc.ClientConnInterface) MongoDBClient {
	return &mongoDBClient{cc}
}

func (c *mongoDBClient) AddMongoDB(ctx context.Context, in *AddMongoDBRequest, opts ...grpc.CallOption) (*AddMongoDBResponse, error) {
	out := new(AddMongoDBResponse)
	err := c.cc.Invoke(ctx, "/management.MongoDB/AddMongoDB", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MongoDBServer is the server API for MongoDB service.
// All implementations must embed UnimplementedMongoDBServer
// for forward compatibility
type MongoDBServer interface {
	// AddMongoDB adds MongoDB Service and starts several Agents.
	// It automatically adds a service to inventory, which is running on provided "node_id",
	// then adds "mongodb_exporter", and "qan_mongodb_profiler" agents
	// with provided "pmm_agent_id" and other parameters.
	AddMongoDB(context.Context, *AddMongoDBRequest) (*AddMongoDBResponse, error)
	mustEmbedUnimplementedMongoDBServer()
}

// UnimplementedMongoDBServer must be embedded to have forward compatible implementations.
type UnimplementedMongoDBServer struct {
}

func (UnimplementedMongoDBServer) AddMongoDB(context.Context, *AddMongoDBRequest) (*AddMongoDBResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMongoDB not implemented")
}
func (UnimplementedMongoDBServer) mustEmbedUnimplementedMongoDBServer() {}

// UnsafeMongoDBServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MongoDBServer will
// result in compilation errors.
type UnsafeMongoDBServer interface {
	mustEmbedUnimplementedMongoDBServer()
}

func RegisterMongoDBServer(s grpc.ServiceRegistrar, srv MongoDBServer) {
	s.RegisterService(&MongoDB_ServiceDesc, srv)
}

func _MongoDB_AddMongoDB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMongoDBRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongoDBServer).AddMongoDB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.MongoDB/AddMongoDB",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongoDBServer).AddMongoDB(ctx, req.(*AddMongoDBRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MongoDB_ServiceDesc is the grpc.ServiceDesc for MongoDB service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MongoDB_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "management.MongoDB",
	HandlerType: (*MongoDBServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddMongoDB",
			Handler:    _MongoDB_AddMongoDB_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "managementpb/mongodb.proto",
}
