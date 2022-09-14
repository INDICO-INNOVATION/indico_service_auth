// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: protos/mfa.proto

package mfa

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

// MFAServiceClient is the client API for MFAService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MFAServiceClient interface {
	GenerateTOTPToken(ctx context.Context, in *GenerateTOTPTokenRequest, opts ...grpc.CallOption) (*GenerateTOTPTokenResponse, error)
	ValidateTOTPToken(ctx context.Context, in *ValidateTOTPTokenRequest, opts ...grpc.CallOption) (*ValidateTOTPTokenResponse, error)
	GenerateSecretKey(ctx context.Context, in *GenerateTOTPTokenRequest, opts ...grpc.CallOption) (*TOTPSecretResponse, error)
}

type mFAServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMFAServiceClient(cc grpc.ClientConnInterface) MFAServiceClient {
	return &mFAServiceClient{cc}
}

func (c *mFAServiceClient) GenerateTOTPToken(ctx context.Context, in *GenerateTOTPTokenRequest, opts ...grpc.CallOption) (*GenerateTOTPTokenResponse, error) {
	out := new(GenerateTOTPTokenResponse)
	err := c.cc.Invoke(ctx, "/MFAService/GenerateTOTPToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mFAServiceClient) ValidateTOTPToken(ctx context.Context, in *ValidateTOTPTokenRequest, opts ...grpc.CallOption) (*ValidateTOTPTokenResponse, error) {
	out := new(ValidateTOTPTokenResponse)
	err := c.cc.Invoke(ctx, "/MFAService/ValidateTOTPToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mFAServiceClient) GenerateSecretKey(ctx context.Context, in *GenerateTOTPTokenRequest, opts ...grpc.CallOption) (*TOTPSecretResponse, error) {
	out := new(TOTPSecretResponse)
	err := c.cc.Invoke(ctx, "/MFAService/GenerateSecretKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MFAServiceServer is the server API for MFAService service.
// All implementations must embed UnimplementedMFAServiceServer
// for forward compatibility
type MFAServiceServer interface {
	GenerateTOTPToken(context.Context, *GenerateTOTPTokenRequest) (*GenerateTOTPTokenResponse, error)
	ValidateTOTPToken(context.Context, *ValidateTOTPTokenRequest) (*ValidateTOTPTokenResponse, error)
	GenerateSecretKey(context.Context, *GenerateTOTPTokenRequest) (*TOTPSecretResponse, error)
	mustEmbedUnimplementedMFAServiceServer()
}

// UnimplementedMFAServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMFAServiceServer struct {
}

func (UnimplementedMFAServiceServer) GenerateTOTPToken(context.Context, *GenerateTOTPTokenRequest) (*GenerateTOTPTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateTOTPToken not implemented")
}
func (UnimplementedMFAServiceServer) ValidateTOTPToken(context.Context, *ValidateTOTPTokenRequest) (*ValidateTOTPTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateTOTPToken not implemented")
}
func (UnimplementedMFAServiceServer) GenerateSecretKey(context.Context, *GenerateTOTPTokenRequest) (*TOTPSecretResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateSecretKey not implemented")
}
func (UnimplementedMFAServiceServer) mustEmbedUnimplementedMFAServiceServer() {}

// UnsafeMFAServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MFAServiceServer will
// result in compilation errors.
type UnsafeMFAServiceServer interface {
	mustEmbedUnimplementedMFAServiceServer()
}

func RegisterMFAServiceServer(s grpc.ServiceRegistrar, srv MFAServiceServer) {
	s.RegisterService(&MFAService_ServiceDesc, srv)
}

func _MFAService_GenerateTOTPToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateTOTPTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MFAServiceServer).GenerateTOTPToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MFAService/GenerateTOTPToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MFAServiceServer).GenerateTOTPToken(ctx, req.(*GenerateTOTPTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MFAService_ValidateTOTPToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateTOTPTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MFAServiceServer).ValidateTOTPToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MFAService/ValidateTOTPToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MFAServiceServer).ValidateTOTPToken(ctx, req.(*ValidateTOTPTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MFAService_GenerateSecretKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateTOTPTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MFAServiceServer).GenerateSecretKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MFAService/GenerateSecretKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MFAServiceServer).GenerateSecretKey(ctx, req.(*GenerateTOTPTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MFAService_ServiceDesc is the grpc.ServiceDesc for MFAService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MFAService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "MFAService",
	HandlerType: (*MFAServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateTOTPToken",
			Handler:    _MFAService_GenerateTOTPToken_Handler,
		},
		{
			MethodName: "ValidateTOTPToken",
			Handler:    _MFAService_ValidateTOTPToken_Handler,
		},
		{
			MethodName: "GenerateSecretKey",
			Handler:    _MFAService_GenerateSecretKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/mfa.proto",
}