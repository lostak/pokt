// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/pokt/tx.proto

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

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	CreatePortfolio(ctx context.Context, in *MsgCreatePortfolio, opts ...grpc.CallOption) (*MsgCreatePortfolioResponse, error)
	CreateAccount(ctx context.Context, in *MsgCreateAccount, opts ...grpc.CallOption) (*MsgCreateAccountResponse, error)
	CreateChain(ctx context.Context, in *MsgCreateChain, opts ...grpc.CallOption) (*MsgCreateChainResponse, error)
	CreateToken(ctx context.Context, in *MsgCreateToken, opts ...grpc.CallOption) (*MsgCreateTokenResponse, error)
	CreateAmount(ctx context.Context, in *MsgCreateAmount, opts ...grpc.CallOption) (*MsgCreateAmountResponse, error)
	UpdateCoinGeckoId(ctx context.Context, in *MsgUpdateCoinGeckoId, opts ...grpc.CallOption) (*MsgUpdateCoinGeckoIdResponse, error)
	ClearPortfolio(ctx context.Context, in *MsgClearPortfolio, opts ...grpc.CallOption) (*MsgClearPortfolioResponse, error)
	ClearAccount(ctx context.Context, in *MsgClearAccount, opts ...grpc.CallOption) (*MsgClearAccountResponse, error)
	ClearChain(ctx context.Context, in *MsgClearChain, opts ...grpc.CallOption) (*MsgClearChainResponse, error)
	ClearToken(ctx context.Context, in *MsgClearToken, opts ...grpc.CallOption) (*MsgClearTokenResponse, error)
	DeletePortfolio(ctx context.Context, in *MsgDeletePortfolio, opts ...grpc.CallOption) (*MsgDeletePortfolioResponse, error)
	DeleteAccount(ctx context.Context, in *MsgDeleteAccount, opts ...grpc.CallOption) (*MsgDeleteAccountResponse, error)
	DeleteChain(ctx context.Context, in *MsgDeleteChain, opts ...grpc.CallOption) (*MsgDeleteChainResponse, error)
	DeleteToken(ctx context.Context, in *MsgDeleteToken, opts ...grpc.CallOption) (*MsgDeleteTokenResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreatePortfolio(ctx context.Context, in *MsgCreatePortfolio, opts ...grpc.CallOption) (*MsgCreatePortfolioResponse, error) {
	out := new(MsgCreatePortfolioResponse)
	err := c.cc.Invoke(ctx, "/pokt.keeper.Msg/CreatePortfolio", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreateAccount(ctx context.Context, in *MsgCreateAccount, opts ...grpc.CallOption) (*MsgCreateAccountResponse, error) {
	out := new(MsgCreateAccountResponse)
	err := c.cc.Invoke(ctx, "/pokt.keeper.Msg/CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreateChain(ctx context.Context, in *MsgCreateChain, opts ...grpc.CallOption) (*MsgCreateChainResponse, error) {
	out := new(MsgCreateChainResponse)
	err := c.cc.Invoke(ctx, "/pokt.keeper.Msg/CreateChain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreateToken(ctx context.Context, in *MsgCreateToken, opts ...grpc.CallOption) (*MsgCreateTokenResponse, error) {
	out := new(MsgCreateTokenResponse)
	err := c.cc.Invoke(ctx, "/pokt.keeper.Msg/CreateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreateAmount(ctx context.Context, in *MsgCreateAmount, opts ...grpc.CallOption) (*MsgCreateAmountResponse, error) {
	out := new(MsgCreateAmountResponse)
	err := c.cc.Invoke(ctx, "/pokt.keeper.Msg/CreateAmount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateCoinGeckoId(ctx context.Context, in *MsgUpdateCoinGeckoId, opts ...grpc.CallOption) (*MsgUpdateCoinGeckoIdResponse, error) {
	out := new(MsgUpdateCoinGeckoIdResponse)
	err := c.cc.Invoke(ctx, "/pokt.keeper.Msg/UpdateCoinGeckoId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ClearPortfolio(ctx context.Context, in *MsgClearPortfolio, opts ...grpc.CallOption) (*MsgClearPortfolioResponse, error) {
	out := new(MsgClearPortfolioResponse)
	err := c.cc.Invoke(ctx, "/pokt.keeper.Msg/ClearPortfolio", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ClearAccount(ctx context.Context, in *MsgClearAccount, opts ...grpc.CallOption) (*MsgClearAccountResponse, error) {
	out := new(MsgClearAccountResponse)
	err := c.cc.Invoke(ctx, "/pokt.keeper.Msg/ClearAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ClearChain(ctx context.Context, in *MsgClearChain, opts ...grpc.CallOption) (*MsgClearChainResponse, error) {
	out := new(MsgClearChainResponse)
	err := c.cc.Invoke(ctx, "/pokt.keeper.Msg/ClearChain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ClearToken(ctx context.Context, in *MsgClearToken, opts ...grpc.CallOption) (*MsgClearTokenResponse, error) {
	out := new(MsgClearTokenResponse)
	err := c.cc.Invoke(ctx, "/pokt.keeper.Msg/ClearToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeletePortfolio(ctx context.Context, in *MsgDeletePortfolio, opts ...grpc.CallOption) (*MsgDeletePortfolioResponse, error) {
	out := new(MsgDeletePortfolioResponse)
	err := c.cc.Invoke(ctx, "/pokt.keeper.Msg/DeletePortfolio", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeleteAccount(ctx context.Context, in *MsgDeleteAccount, opts ...grpc.CallOption) (*MsgDeleteAccountResponse, error) {
	out := new(MsgDeleteAccountResponse)
	err := c.cc.Invoke(ctx, "/pokt.keeper.Msg/DeleteAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeleteChain(ctx context.Context, in *MsgDeleteChain, opts ...grpc.CallOption) (*MsgDeleteChainResponse, error) {
	out := new(MsgDeleteChainResponse)
	err := c.cc.Invoke(ctx, "/pokt.keeper.Msg/DeleteChain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeleteToken(ctx context.Context, in *MsgDeleteToken, opts ...grpc.CallOption) (*MsgDeleteTokenResponse, error) {
	out := new(MsgDeleteTokenResponse)
	err := c.cc.Invoke(ctx, "/pokt.keeper.Msg/DeleteToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	CreatePortfolio(context.Context, *MsgCreatePortfolio) (*MsgCreatePortfolioResponse, error)
	CreateAccount(context.Context, *MsgCreateAccount) (*MsgCreateAccountResponse, error)
	CreateChain(context.Context, *MsgCreateChain) (*MsgCreateChainResponse, error)
	CreateToken(context.Context, *MsgCreateToken) (*MsgCreateTokenResponse, error)
	CreateAmount(context.Context, *MsgCreateAmount) (*MsgCreateAmountResponse, error)
	UpdateCoinGeckoId(context.Context, *MsgUpdateCoinGeckoId) (*MsgUpdateCoinGeckoIdResponse, error)
	ClearPortfolio(context.Context, *MsgClearPortfolio) (*MsgClearPortfolioResponse, error)
	ClearAccount(context.Context, *MsgClearAccount) (*MsgClearAccountResponse, error)
	ClearChain(context.Context, *MsgClearChain) (*MsgClearChainResponse, error)
	ClearToken(context.Context, *MsgClearToken) (*MsgClearTokenResponse, error)
	DeletePortfolio(context.Context, *MsgDeletePortfolio) (*MsgDeletePortfolioResponse, error)
	DeleteAccount(context.Context, *MsgDeleteAccount) (*MsgDeleteAccountResponse, error)
	DeleteChain(context.Context, *MsgDeleteChain) (*MsgDeleteChainResponse, error)
	DeleteToken(context.Context, *MsgDeleteToken) (*MsgDeleteTokenResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) CreatePortfolio(context.Context, *MsgCreatePortfolio) (*MsgCreatePortfolioResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePortfolio not implemented")
}
func (UnimplementedMsgServer) CreateAccount(context.Context, *MsgCreateAccount) (*MsgCreateAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedMsgServer) CreateChain(context.Context, *MsgCreateChain) (*MsgCreateChainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChain not implemented")
}
func (UnimplementedMsgServer) CreateToken(context.Context, *MsgCreateToken) (*MsgCreateTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateToken not implemented")
}
func (UnimplementedMsgServer) CreateAmount(context.Context, *MsgCreateAmount) (*MsgCreateAmountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAmount not implemented")
}
func (UnimplementedMsgServer) UpdateCoinGeckoId(context.Context, *MsgUpdateCoinGeckoId) (*MsgUpdateCoinGeckoIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCoinGeckoId not implemented")
}
func (UnimplementedMsgServer) ClearPortfolio(context.Context, *MsgClearPortfolio) (*MsgClearPortfolioResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearPortfolio not implemented")
}
func (UnimplementedMsgServer) ClearAccount(context.Context, *MsgClearAccount) (*MsgClearAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearAccount not implemented")
}
func (UnimplementedMsgServer) ClearChain(context.Context, *MsgClearChain) (*MsgClearChainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearChain not implemented")
}
func (UnimplementedMsgServer) ClearToken(context.Context, *MsgClearToken) (*MsgClearTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearToken not implemented")
}
func (UnimplementedMsgServer) DeletePortfolio(context.Context, *MsgDeletePortfolio) (*MsgDeletePortfolioResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePortfolio not implemented")
}
func (UnimplementedMsgServer) DeleteAccount(context.Context, *MsgDeleteAccount) (*MsgDeleteAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAccount not implemented")
}
func (UnimplementedMsgServer) DeleteChain(context.Context, *MsgDeleteChain) (*MsgDeleteChainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteChain not implemented")
}
func (UnimplementedMsgServer) DeleteToken(context.Context, *MsgDeleteToken) (*MsgDeleteTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteToken not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_CreatePortfolio_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreatePortfolio)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreatePortfolio(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pokt.keeper.Msg/CreatePortfolio",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreatePortfolio(ctx, req.(*MsgCreatePortfolio))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pokt.keeper.Msg/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateAccount(ctx, req.(*MsgCreateAccount))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreateChain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateChain)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateChain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pokt.keeper.Msg/CreateChain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateChain(ctx, req.(*MsgCreateChain))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pokt.keeper.Msg/CreateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateToken(ctx, req.(*MsgCreateToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreateAmount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateAmount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateAmount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pokt.keeper.Msg/CreateAmount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateAmount(ctx, req.(*MsgCreateAmount))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateCoinGeckoId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateCoinGeckoId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateCoinGeckoId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pokt.keeper.Msg/UpdateCoinGeckoId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateCoinGeckoId(ctx, req.(*MsgUpdateCoinGeckoId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ClearPortfolio_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgClearPortfolio)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ClearPortfolio(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pokt.keeper.Msg/ClearPortfolio",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ClearPortfolio(ctx, req.(*MsgClearPortfolio))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ClearAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgClearAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ClearAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pokt.keeper.Msg/ClearAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ClearAccount(ctx, req.(*MsgClearAccount))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ClearChain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgClearChain)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ClearChain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pokt.keeper.Msg/ClearChain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ClearChain(ctx, req.(*MsgClearChain))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ClearToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgClearToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ClearToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pokt.keeper.Msg/ClearToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ClearToken(ctx, req.(*MsgClearToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeletePortfolio_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeletePortfolio)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeletePortfolio(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pokt.keeper.Msg/DeletePortfolio",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeletePortfolio(ctx, req.(*MsgDeletePortfolio))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeleteAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeleteAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeleteAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pokt.keeper.Msg/DeleteAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeleteAccount(ctx, req.(*MsgDeleteAccount))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeleteChain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeleteChain)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeleteChain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pokt.keeper.Msg/DeleteChain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeleteChain(ctx, req.(*MsgDeleteChain))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeleteToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeleteToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeleteToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pokt.keeper.Msg/DeleteToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeleteToken(ctx, req.(*MsgDeleteToken))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pokt.keeper.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePortfolio",
			Handler:    _Msg_CreatePortfolio_Handler,
		},
		{
			MethodName: "CreateAccount",
			Handler:    _Msg_CreateAccount_Handler,
		},
		{
			MethodName: "CreateChain",
			Handler:    _Msg_CreateChain_Handler,
		},
		{
			MethodName: "CreateToken",
			Handler:    _Msg_CreateToken_Handler,
		},
		{
			MethodName: "CreateAmount",
			Handler:    _Msg_CreateAmount_Handler,
		},
		{
			MethodName: "UpdateCoinGeckoId",
			Handler:    _Msg_UpdateCoinGeckoId_Handler,
		},
		{
			MethodName: "ClearPortfolio",
			Handler:    _Msg_ClearPortfolio_Handler,
		},
		{
			MethodName: "ClearAccount",
			Handler:    _Msg_ClearAccount_Handler,
		},
		{
			MethodName: "ClearChain",
			Handler:    _Msg_ClearChain_Handler,
		},
		{
			MethodName: "ClearToken",
			Handler:    _Msg_ClearToken_Handler,
		},
		{
			MethodName: "DeletePortfolio",
			Handler:    _Msg_DeletePortfolio_Handler,
		},
		{
			MethodName: "DeleteAccount",
			Handler:    _Msg_DeleteAccount_Handler,
		},
		{
			MethodName: "DeleteChain",
			Handler:    _Msg_DeleteChain_Handler,
		},
		{
			MethodName: "DeleteToken",
			Handler:    _Msg_DeleteToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/pokt/tx.proto",
}
