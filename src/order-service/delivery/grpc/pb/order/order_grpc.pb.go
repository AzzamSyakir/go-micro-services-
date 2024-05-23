// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: order.proto

package order

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

const (
	OrderService_GetOrderById_FullMethodName  = "/go_micro_services.OrderService/GetOrderById"
	OrderService_Order_FullMethodName         = "/go_micro_services.OrderService/Order"
	OrderService_ListOrders_FullMethodName    = "/go_micro_services.OrderService/ListOrders"
	OrderService_OrderProducts_FullMethodName = "/go_micro_services.OrderService/OrderProducts"
)

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderServiceClient interface {
	GetOrderById(ctx context.Context, in *ById, opts ...grpc.CallOption) (*OrderResponse, error)
	Order(ctx context.Context, in *Create, opts ...grpc.CallOption) (*OrderResponse, error)
	ListOrders(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*OrderResponseRepeated, error)
	OrderProducts(ctx context.Context, in *OrderProductRequest, opts ...grpc.CallOption) (*OrderProductResponse, error)
}

type orderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderServiceClient(cc grpc.ClientConnInterface) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) GetOrderById(ctx context.Context, in *ById, opts ...grpc.CallOption) (*OrderResponse, error) {
	out := new(OrderResponse)
	err := c.cc.Invoke(ctx, OrderService_GetOrderById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) Order(ctx context.Context, in *Create, opts ...grpc.CallOption) (*OrderResponse, error) {
	out := new(OrderResponse)
	err := c.cc.Invoke(ctx, OrderService_Order_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) ListOrders(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*OrderResponseRepeated, error) {
	out := new(OrderResponseRepeated)
	err := c.cc.Invoke(ctx, OrderService_ListOrders_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) OrderProducts(ctx context.Context, in *OrderProductRequest, opts ...grpc.CallOption) (*OrderProductResponse, error) {
	out := new(OrderProductResponse)
	err := c.cc.Invoke(ctx, OrderService_OrderProducts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServiceServer is the server API for OrderService service.
// All implementations must embed UnimplementedOrderServiceServer
// for forward compatibility
type OrderServiceServer interface {
	GetOrderById(context.Context, *ById) (*OrderResponse, error)
	Order(context.Context, *Create) (*OrderResponse, error)
	ListOrders(context.Context, *Empty) (*OrderResponseRepeated, error)
	OrderProducts(context.Context, *OrderProductRequest) (*OrderProductResponse, error)
	mustEmbedUnimplementedOrderServiceServer()
}

// UnimplementedOrderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrderServiceServer struct {
}

func (UnimplementedOrderServiceServer) GetOrderById(context.Context, *ById) (*OrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderById not implemented")
}
func (UnimplementedOrderServiceServer) Order(context.Context, *Create) (*OrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Order not implemented")
}
func (UnimplementedOrderServiceServer) ListOrders(context.Context, *Empty) (*OrderResponseRepeated, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrders not implemented")
}
func (UnimplementedOrderServiceServer) OrderProducts(context.Context, *OrderProductRequest) (*OrderProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderProducts not implemented")
}
func (UnimplementedOrderServiceServer) mustEmbedUnimplementedOrderServiceServer() {}

// UnsafeOrderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServiceServer will
// result in compilation errors.
type UnsafeOrderServiceServer interface {
	mustEmbedUnimplementedOrderServiceServer()
}

func RegisterOrderServiceServer(s grpc.ServiceRegistrar, srv OrderServiceServer) {
	s.RegisterService(&OrderService_ServiceDesc, srv)
}

func _OrderService_GetOrderById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetOrderById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_GetOrderById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetOrderById(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_Order_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Create)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).Order(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_Order_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).Order(ctx, req.(*Create))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_ListOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).ListOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_ListOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).ListOrders(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_OrderProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).OrderProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_OrderProducts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).OrderProducts(ctx, req.(*OrderProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderService_ServiceDesc is the grpc.ServiceDesc for OrderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "go_micro_services.OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOrderById",
			Handler:    _OrderService_GetOrderById_Handler,
		},
		{
			MethodName: "Order",
			Handler:    _OrderService_Order_Handler,
		},
		{
			MethodName: "ListOrders",
			Handler:    _OrderService_ListOrders_Handler,
		},
		{
			MethodName: "OrderProducts",
			Handler:    _OrderService_OrderProducts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order.proto",
}
