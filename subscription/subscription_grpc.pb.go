// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: subscription/subscription.proto

package subscription

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

// MembershipServiceClient is the client API for MembershipService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MembershipServiceClient interface {
	GetMembership(ctx context.Context, in *Membership, opts ...grpc.CallOption) (*Membership, error)
}

type membershipServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMembershipServiceClient(cc grpc.ClientConnInterface) MembershipServiceClient {
	return &membershipServiceClient{cc}
}

func (c *membershipServiceClient) GetMembership(ctx context.Context, in *Membership, opts ...grpc.CallOption) (*Membership, error) {
	out := new(Membership)
	err := c.cc.Invoke(ctx, "/subscription.MembershipService/GetMembership", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MembershipServiceServer is the server API for MembershipService service.
// All implementations must embed UnimplementedMembershipServiceServer
// for forward compatibility
type MembershipServiceServer interface {
	GetMembership(context.Context, *Membership) (*Membership, error)
	mustEmbedUnimplementedMembershipServiceServer()
}

// UnimplementedMembershipServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMembershipServiceServer struct {
}

func (UnimplementedMembershipServiceServer) GetMembership(context.Context, *Membership) (*Membership, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMembership not implemented")
}
func (UnimplementedMembershipServiceServer) mustEmbedUnimplementedMembershipServiceServer() {}

// UnsafeMembershipServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MembershipServiceServer will
// result in compilation errors.
type UnsafeMembershipServiceServer interface {
	mustEmbedUnimplementedMembershipServiceServer()
}

func RegisterMembershipServiceServer(s grpc.ServiceRegistrar, srv MembershipServiceServer) {
	s.RegisterService(&MembershipService_ServiceDesc, srv)
}

func _MembershipService_GetMembership_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Membership)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MembershipServiceServer).GetMembership(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/subscription.MembershipService/GetMembership",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MembershipServiceServer).GetMembership(ctx, req.(*Membership))
	}
	return interceptor(ctx, in, info, handler)
}

// MembershipService_ServiceDesc is the grpc.ServiceDesc for MembershipService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MembershipService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "subscription.MembershipService",
	HandlerType: (*MembershipServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMembership",
			Handler:    _MembershipService_GetMembership_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "subscription/subscription.proto",
}
