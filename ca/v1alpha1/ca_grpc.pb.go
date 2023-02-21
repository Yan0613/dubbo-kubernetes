// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: v1alpha1/ca.proto

package v1alpha1

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

// DubboCertificateServiceClient is the client API for DubboCertificateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DubboCertificateServiceClient interface {
	// Using provided CSR, returns a signed certificate.
	CreateCertificate(ctx context.Context, in *DubboCertificateRequest, opts ...grpc.CallOption) (*DubboCertificateResponse, error)
}

type dubboCertificateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDubboCertificateServiceClient(cc grpc.ClientConnInterface) DubboCertificateServiceClient {
	return &dubboCertificateServiceClient{cc}
}

func (c *dubboCertificateServiceClient) CreateCertificate(ctx context.Context, in *DubboCertificateRequest, opts ...grpc.CallOption) (*DubboCertificateResponse, error) {
	out := new(DubboCertificateResponse)
	err := c.cc.Invoke(ctx, "/org.apache.dubbo.auth.v1alpha1.DubboCertificateService/CreateCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DubboCertificateServiceServer is the server API for DubboCertificateService service.
// All implementations must embed UnimplementedDubboCertificateServiceServer
// for forward compatibility
type DubboCertificateServiceServer interface {
	// Using provided CSR, returns a signed certificate.
	CreateCertificate(context.Context, *DubboCertificateRequest) (*DubboCertificateResponse, error)
	mustEmbedUnimplementedDubboCertificateServiceServer()
}

// UnimplementedDubboCertificateServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDubboCertificateServiceServer struct {
}

func (UnimplementedDubboCertificateServiceServer) CreateCertificate(context.Context, *DubboCertificateRequest) (*DubboCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCertificate not implemented")
}
func (UnimplementedDubboCertificateServiceServer) mustEmbedUnimplementedDubboCertificateServiceServer() {
}

// UnsafeDubboCertificateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DubboCertificateServiceServer will
// result in compilation errors.
type UnsafeDubboCertificateServiceServer interface {
	mustEmbedUnimplementedDubboCertificateServiceServer()
}

func RegisterDubboCertificateServiceServer(s grpc.ServiceRegistrar, srv DubboCertificateServiceServer) {
	s.RegisterService(&DubboCertificateService_ServiceDesc, srv)
}

func _DubboCertificateService_CreateCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DubboCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DubboCertificateServiceServer).CreateCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/org.apache.dubbo.auth.v1alpha1.DubboCertificateService/CreateCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DubboCertificateServiceServer).CreateCertificate(ctx, req.(*DubboCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DubboCertificateService_ServiceDesc is the grpc.ServiceDesc for DubboCertificateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DubboCertificateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "org.apache.dubbo.auth.v1alpha1.DubboCertificateService",
	HandlerType: (*DubboCertificateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCertificate",
			Handler:    _DubboCertificateService_CreateCertificate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1alpha1/ca.proto",
}
