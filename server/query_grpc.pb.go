// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/pokt/query.proto

package server

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

// QueryServiceClient is the client API for QueryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryServiceClient interface {
	GetPortfolio(ctx context.Context, in *MsgGetPortfolio, opts ...grpc.CallOption) (*MsgGetPortfolioResponse, error)
	// rpc GetAccounts(MsgGetAccounts) returns (stream MsgGetAccountsResponse);
	GetAccount(ctx context.Context, in *MsgGetAccount, opts ...grpc.CallOption) (*MsgGetAccountResponse, error)
	// rpc GetChains(MsgGetChains) returns (stream MsgGetChainsResponse)
	GetChain(ctx context.Context, in *MsgGetChain, opts ...grpc.CallOption) (*MsgGetChainResponse, error)
	// rpc GetTokens(MsgGetTokens) returns (stream MsgGetTokensResponse)
	GetToken(ctx context.Context, in *MsgGetToken, opts ...grpc.CallOption) (*MsgGetTokenResponse, error)
	GetAmounts(ctx context.Context, in *MsgGetAmounts, opts ...grpc.CallOption) (*MsgGetAmountsResponse, error)
}

type queryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryServiceClient(cc grpc.ClientConnInterface) QueryServiceClient {
	return &queryServiceClient{cc}
}

func (c *queryServiceClient) GetPortfolio(ctx context.Context, in *MsgGetPortfolio, opts ...grpc.CallOption) (*MsgGetPortfolioResponse, error) {
	out := new(MsgGetPortfolioResponse)
	err := c.cc.Invoke(ctx, "/pokt.server.QueryService/GetPortfolio", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryServiceClient) GetAccount(ctx context.Context, in *MsgGetAccount, opts ...grpc.CallOption) (*MsgGetAccountResponse, error) {
	out := new(MsgGetAccountResponse)
	err := c.cc.Invoke(ctx, "/pokt.server.QueryService/GetAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryServiceClient) GetChain(ctx context.Context, in *MsgGetChain, opts ...grpc.CallOption) (*MsgGetChainResponse, error) {
	out := new(MsgGetChainResponse)
	err := c.cc.Invoke(ctx, "/pokt.server.QueryService/GetChain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryServiceClient) GetToken(ctx context.Context, in *MsgGetToken, opts ...grpc.CallOption) (*MsgGetTokenResponse, error) {
	out := new(MsgGetTokenResponse)
	err := c.cc.Invoke(ctx, "/pokt.server.QueryService/GetToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryServiceClient) GetAmounts(ctx context.Context, in *MsgGetAmounts, opts ...grpc.CallOption) (*MsgGetAmountsResponse, error) {
	out := new(MsgGetAmountsResponse)
	err := c.cc.Invoke(ctx, "/pokt.server.QueryService/GetAmounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServiceServer is the server API for QueryService service.
// All implementations must embed UnimplementedQueryServiceServer
// for forward compatibility
type QueryServiceServer interface {
	GetPortfolio(context.Context, *MsgGetPortfolio) (*MsgGetPortfolioResponse, error)
	// rpc GetAccounts(MsgGetAccounts) returns (stream MsgGetAccountsResponse);
	GetAccount(context.Context, *MsgGetAccount) (*MsgGetAccountResponse, error)
	// rpc GetChains(MsgGetChains) returns (stream MsgGetChainsResponse)
	GetChain(context.Context, *MsgGetChain) (*MsgGetChainResponse, error)
	// rpc GetTokens(MsgGetTokens) returns (stream MsgGetTokensResponse)
	GetToken(context.Context, *MsgGetToken) (*MsgGetTokenResponse, error)
	GetAmounts(context.Context, *MsgGetAmounts) (*MsgGetAmountsResponse, error)
	mustEmbedUnimplementedQueryServiceServer()
}

// UnimplementedQueryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServiceServer struct {
}

func (UnimplementedQueryServiceServer) GetPortfolio(context.Context, *MsgGetPortfolio) (*MsgGetPortfolioResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPortfolio not implemented")
}
func (UnimplementedQueryServiceServer) GetAccount(context.Context, *MsgGetAccount) (*MsgGetAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccount not implemented")
}
func (UnimplementedQueryServiceServer) GetChain(context.Context, *MsgGetChain) (*MsgGetChainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChain not implemented")
}
func (UnimplementedQueryServiceServer) GetToken(context.Context, *MsgGetToken) (*MsgGetTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetToken not implemented")
}
func (UnimplementedQueryServiceServer) GetAmounts(context.Context, *MsgGetAmounts) (*MsgGetAmountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAmounts not implemented")
}
func (UnimplementedQueryServiceServer) mustEmbedUnimplementedQueryServiceServer() {}

// UnsafeQueryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServiceServer will
// result in compilation errors.
type UnsafeQueryServiceServer interface {
	mustEmbedUnimplementedQueryServiceServer()
}

func RegisterQueryServiceServer(s grpc.ServiceRegistrar, srv QueryServiceServer) {
	s.RegisterService(&QueryService_ServiceDesc, srv)
}

func _QueryService_GetPortfolio_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgGetPortfolio)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).GetPortfolio(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pokt.server.QueryService/GetPortfolio",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).GetPortfolio(ctx, req.(*MsgGetPortfolio))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryService_GetAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgGetAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).GetAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pokt.server.QueryService/GetAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).GetAccount(ctx, req.(*MsgGetAccount))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryService_GetChain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgGetChain)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).GetChain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pokt.server.QueryService/GetChain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).GetChain(ctx, req.(*MsgGetChain))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryService_GetToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgGetToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).GetToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pokt.server.QueryService/GetToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).GetToken(ctx, req.(*MsgGetToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryService_GetAmounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgGetAmounts)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).GetAmounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pokt.server.QueryService/GetAmounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).GetAmounts(ctx, req.(*MsgGetAmounts))
	}
	return interceptor(ctx, in, info, handler)
}

// QueryService_ServiceDesc is the grpc.ServiceDesc for QueryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QueryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pokt.server.QueryService",
	HandlerType: (*QueryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPortfolio",
			Handler:    _QueryService_GetPortfolio_Handler,
		},
		{
			MethodName: "GetAccount",
			Handler:    _QueryService_GetAccount_Handler,
		},
		{
			MethodName: "GetChain",
			Handler:    _QueryService_GetChain_Handler,
		},
		{
			MethodName: "GetToken",
			Handler:    _QueryService_GetToken_Handler,
		},
		{
			MethodName: "GetAmounts",
			Handler:    _QueryService_GetAmounts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/pokt/query.proto",
}