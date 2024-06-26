// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.6.1
// source: rpc/infofile.proto

package rpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	InfoFileService_LoadFile_FullMethodName = "/InfoFileService/LoadFile"
)

// InfoFileServiceClient is the client API for InfoFileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InfoFileServiceClient interface {
	LoadFile(ctx context.Context, in *InfoFileRequest, opts ...grpc.CallOption) (*InfoFileResponse, error)
}

type infoFileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInfoFileServiceClient(cc grpc.ClientConnInterface) InfoFileServiceClient {
	return &infoFileServiceClient{cc}
}

func (c *infoFileServiceClient) LoadFile(ctx context.Context, in *InfoFileRequest, opts ...grpc.CallOption) (*InfoFileResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(InfoFileResponse)
	err := c.cc.Invoke(ctx, InfoFileService_LoadFile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InfoFileServiceServer is the server API for InfoFileService service.
// All implementations must embed UnimplementedInfoFileServiceServer
// for forward compatibility
type InfoFileServiceServer interface {
	LoadFile(context.Context, *InfoFileRequest) (*InfoFileResponse, error)
	mustEmbedUnimplementedInfoFileServiceServer()
}

// UnimplementedInfoFileServiceServer must be embedded to have forward compatible implementations.
type UnimplementedInfoFileServiceServer struct {
}

func (UnimplementedInfoFileServiceServer) LoadFile(context.Context, *InfoFileRequest) (*InfoFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadFile not implemented")
}
func (UnimplementedInfoFileServiceServer) mustEmbedUnimplementedInfoFileServiceServer() {}

// UnsafeInfoFileServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InfoFileServiceServer will
// result in compilation errors.
type UnsafeInfoFileServiceServer interface {
	mustEmbedUnimplementedInfoFileServiceServer()
}

func RegisterInfoFileServiceServer(s grpc.ServiceRegistrar, srv InfoFileServiceServer) {
	s.RegisterService(&InfoFileService_ServiceDesc, srv)
}

func _InfoFileService_LoadFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoFileServiceServer).LoadFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InfoFileService_LoadFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoFileServiceServer).LoadFile(ctx, req.(*InfoFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InfoFileService_ServiceDesc is the grpc.ServiceDesc for InfoFileService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InfoFileService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "InfoFileService",
	HandlerType: (*InfoFileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoadFile",
			Handler:    _InfoFileService_LoadFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc/infofile.proto",
}
