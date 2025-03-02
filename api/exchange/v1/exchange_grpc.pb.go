// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: exchange/v1/exchange.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Exchange_GetSymbols_FullMethodName = "/exchange.v1.Exchange/GetSymbols"
	Exchange_GetKline_FullMethodName   = "/exchange.v1.Exchange/GetKline"
	Exchange_GetSymbol_FullMethodName  = "/exchange.v1.Exchange/GetSymbol"
)

// ExchangeClient is the client API for Exchange service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExchangeClient interface {
	// 获取所有交易对
	GetSymbols(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SymbolsReply, error)
	// 获取指定交易对的K线数据
	GetKline(ctx context.Context, in *KlineRequest, opts ...grpc.CallOption) (*KlineReply, error)
	// 获取指定交易对的信息
	GetSymbol(ctx context.Context, in *SymbolRequest, opts ...grpc.CallOption) (*SymbolReply, error)
}

type exchangeClient struct {
	cc grpc.ClientConnInterface
}

func NewExchangeClient(cc grpc.ClientConnInterface) ExchangeClient {
	return &exchangeClient{cc}
}

func (c *exchangeClient) GetSymbols(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SymbolsReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SymbolsReply)
	err := c.cc.Invoke(ctx, Exchange_GetSymbols_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeClient) GetKline(ctx context.Context, in *KlineRequest, opts ...grpc.CallOption) (*KlineReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(KlineReply)
	err := c.cc.Invoke(ctx, Exchange_GetKline_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeClient) GetSymbol(ctx context.Context, in *SymbolRequest, opts ...grpc.CallOption) (*SymbolReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SymbolReply)
	err := c.cc.Invoke(ctx, Exchange_GetSymbol_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExchangeServer is the server API for Exchange service.
// All implementations must embed UnimplementedExchangeServer
// for forward compatibility.
type ExchangeServer interface {
	// 获取所有交易对
	GetSymbols(context.Context, *emptypb.Empty) (*SymbolsReply, error)
	// 获取指定交易对的K线数据
	GetKline(context.Context, *KlineRequest) (*KlineReply, error)
	// 获取指定交易对的信息
	GetSymbol(context.Context, *SymbolRequest) (*SymbolReply, error)
	mustEmbedUnimplementedExchangeServer()
}

// UnimplementedExchangeServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedExchangeServer struct{}

func (UnimplementedExchangeServer) GetSymbols(context.Context, *emptypb.Empty) (*SymbolsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSymbols not implemented")
}
func (UnimplementedExchangeServer) GetKline(context.Context, *KlineRequest) (*KlineReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKline not implemented")
}
func (UnimplementedExchangeServer) GetSymbol(context.Context, *SymbolRequest) (*SymbolReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSymbol not implemented")
}
func (UnimplementedExchangeServer) mustEmbedUnimplementedExchangeServer() {}
func (UnimplementedExchangeServer) testEmbeddedByValue()                  {}

// UnsafeExchangeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExchangeServer will
// result in compilation errors.
type UnsafeExchangeServer interface {
	mustEmbedUnimplementedExchangeServer()
}

func RegisterExchangeServer(s grpc.ServiceRegistrar, srv ExchangeServer) {
	// If the following call pancis, it indicates UnimplementedExchangeServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Exchange_ServiceDesc, srv)
}

func _Exchange_GetSymbols_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServer).GetSymbols(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Exchange_GetSymbols_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServer).GetSymbols(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Exchange_GetKline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KlineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServer).GetKline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Exchange_GetKline_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServer).GetKline(ctx, req.(*KlineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Exchange_GetSymbol_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SymbolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServer).GetSymbol(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Exchange_GetSymbol_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServer).GetSymbol(ctx, req.(*SymbolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Exchange_ServiceDesc is the grpc.ServiceDesc for Exchange service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Exchange_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "exchange.v1.Exchange",
	HandlerType: (*ExchangeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSymbols",
			Handler:    _Exchange_GetSymbols_Handler,
		},
		{
			MethodName: "GetKline",
			Handler:    _Exchange_GetKline_Handler,
		},
		{
			MethodName: "GetSymbol",
			Handler:    _Exchange_GetSymbol_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "exchange/v1/exchange.proto",
}
