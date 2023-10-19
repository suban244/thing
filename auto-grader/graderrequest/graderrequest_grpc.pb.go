// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.3
// source: graderrequest.proto

package graderrequest

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

// GraderRequestServiceClient is the client API for GraderRequestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GraderRequestServiceClient interface {
	GradeFile(ctx context.Context, in *File, opts ...grpc.CallOption) (*Status, error)
}

type graderRequestServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGraderRequestServiceClient(cc grpc.ClientConnInterface) GraderRequestServiceClient {
	return &graderRequestServiceClient{cc}
}

func (c *graderRequestServiceClient) GradeFile(ctx context.Context, in *File, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/GraderRequestPackage.GraderRequestService/GradeFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GraderRequestServiceServer is the server API for GraderRequestService service.
// All implementations must embed UnimplementedGraderRequestServiceServer
// for forward compatibility
type GraderRequestServiceServer interface {
	GradeFile(context.Context, *File) (*Status, error)
	mustEmbedUnimplementedGraderRequestServiceServer()
}

// UnimplementedGraderRequestServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGraderRequestServiceServer struct {
}

func (UnimplementedGraderRequestServiceServer) GradeFile(context.Context, *File) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GradeFile not implemented")
}
func (UnimplementedGraderRequestServiceServer) mustEmbedUnimplementedGraderRequestServiceServer() {}

// UnsafeGraderRequestServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GraderRequestServiceServer will
// result in compilation errors.
type UnsafeGraderRequestServiceServer interface {
	mustEmbedUnimplementedGraderRequestServiceServer()
}

func RegisterGraderRequestServiceServer(s grpc.ServiceRegistrar, srv GraderRequestServiceServer) {
	s.RegisterService(&GraderRequestService_ServiceDesc, srv)
}

func _GraderRequestService_GradeFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(File)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GraderRequestServiceServer).GradeFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GraderRequestPackage.GraderRequestService/GradeFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GraderRequestServiceServer).GradeFile(ctx, req.(*File))
	}
	return interceptor(ctx, in, info, handler)
}

// GraderRequestService_ServiceDesc is the grpc.ServiceDesc for GraderRequestService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GraderRequestService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "GraderRequestPackage.GraderRequestService",
	HandlerType: (*GraderRequestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GradeFile",
			Handler:    _GraderRequestService_GradeFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "graderrequest.proto",
}
