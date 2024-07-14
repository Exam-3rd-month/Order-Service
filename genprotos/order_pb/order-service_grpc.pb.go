// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.21.12
// source: order-service/order-service.proto

package order_pb

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
	OrderService_AddDish_FullMethodName                   = "/OrderService/AddDish"
	OrderService_UpdateDish_FullMethodName                = "/OrderService/UpdateDish"
	OrderService_DeleteDish_FullMethodName                = "/OrderService/DeleteDish"
	OrderService_ListDishes_FullMethodName                = "/OrderService/ListDishes"
	OrderService_CreateOrder_FullMethodName               = "/OrderService/CreateOrder"
	OrderService_UpdateOrderStatus_FullMethodName         = "/OrderService/UpdateOrderStatus"
	OrderService_ListOfOrders_FullMethodName              = "/OrderService/ListOfOrders"
	OrderService_GetOrderByKitchenId_FullMethodName       = "/OrderService/GetOrderByKitchenId"
	OrderService_AddReview_FullMethodName                 = "/OrderService/AddReview"
	OrderService_ListReviews_FullMethodName               = "/OrderService/ListReviews"
	OrderService_CreatePayment_FullMethodName             = "/OrderService/CreatePayment"
	OrderService_GetFullInfoAboutOrder_FullMethodName     = "/OrderService/GetFullInfoAboutOrder"
	OrderService_GetDishRecommendations_FullMethodName    = "/OrderService/GetDishRecommendations"
	OrderService_GetKitchenStatistics_FullMethodName      = "/OrderService/GetKitchenStatistics"
	OrderService_GetUserActivity_FullMethodName           = "/OrderService/GetUserActivity"
	OrderService_CreateKitchenWorkingHours_FullMethodName = "/OrderService/CreateKitchenWorkingHours"
	OrderService_UpdateKitchenWorkingHours_FullMethodName = "/OrderService/UpdateKitchenWorkingHours"
	OrderService_CreateDishNutritionInfo_FullMethodName   = "/OrderService/CreateDishNutritionInfo"
	OrderService_UpdateDishNutritionInfo_FullMethodName   = "/OrderService/UpdateDishNutritionInfo"
)

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderServiceClient interface {
	// Dish
	// 1 Done
	AddDish(ctx context.Context, in *AddDishRequest, opts ...grpc.CallOption) (*AddDishResponse, error)
	// 2 Done
	UpdateDish(ctx context.Context, in *UpdateDishRequest, opts ...grpc.CallOption) (*UpdateDishResponse, error)
	// 3 Done
	DeleteDish(ctx context.Context, in *DeleteDishRequest, opts ...grpc.CallOption) (*DeleteDishResponse, error)
	// 4 Done
	ListDishes(ctx context.Context, in *ListDishesRequest, opts ...grpc.CallOption) (*ListDishesResponse, error)
	// Order
	// 5 Done
	CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error)
	// 6 Done
	UpdateOrderStatus(ctx context.Context, in *UpdateOrderStatusRequest, opts ...grpc.CallOption) (*UpdateOrderStatusResponse, error)
	// 7 Done
	ListOfOrders(ctx context.Context, in *ListOfOrdersRequest, opts ...grpc.CallOption) (*ListOfOrdersResponse, error)
	// 8 Done
	GetOrderByKitchenId(ctx context.Context, in *GetOrderByKitchenIdRequest, opts ...grpc.CallOption) (*GetOrderByKitchenIdResponse, error)
	// 9 Done
	AddReview(ctx context.Context, in *AddReviewRequest, opts ...grpc.CallOption) (*AddReviewResponse, error)
	// 10 Done
	ListReviews(ctx context.Context, in *ListReviewsRequest, opts ...grpc.CallOption) (*ListReviewsResponse, error)
	// 11
	CreatePayment(ctx context.Context, in *CreatePaymentRequest, opts ...grpc.CallOption) (*CreatePaymentResponse, error)
	// 12
	GetFullInfoAboutOrder(ctx context.Context, in *GetFullInfoAboutOrderRequest, opts ...grpc.CallOption) (*GetFullInfoAboutOrderResponse, error)
	// Qo'shimchalar
	// 1
	GetDishRecommendations(ctx context.Context, in *GetDishRecommendationsRequest, opts ...grpc.CallOption) (*GetDishRecommendationsResponse, error)
	// 2
	GetKitchenStatistics(ctx context.Context, in *GetKitchenStatisticsRequest, opts ...grpc.CallOption) (*GetKitchenStatisticsResponse, error)
	// 3
	GetUserActivity(ctx context.Context, in *GetUserActivityRequest, opts ...grpc.CallOption) (*GetUserActivityResponse, error)
	// 4
	CreateKitchenWorkingHours(ctx context.Context, in *CreateKitchenWorkingHoursRequest, opts ...grpc.CallOption) (*CreateKitchenWorkingHoursResponse, error)
	// 5
	UpdateKitchenWorkingHours(ctx context.Context, in *UpdateKitchenWorkingHoursRequest, opts ...grpc.CallOption) (*UpdateKitchenWorkingHoursResponse, error)
	// 6
	CreateDishNutritionInfo(ctx context.Context, in *CreateDishNutritionInfoRequest, opts ...grpc.CallOption) (*CreateDishNutritionInfoResponse, error)
	// 7
	UpdateDishNutritionInfo(ctx context.Context, in *UpdateDishNutritionInfoRequest, opts ...grpc.CallOption) (*UpdateDishNutritionInfoResponse, error)
}

type orderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderServiceClient(cc grpc.ClientConnInterface) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) AddDish(ctx context.Context, in *AddDishRequest, opts ...grpc.CallOption) (*AddDishResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddDishResponse)
	err := c.cc.Invoke(ctx, OrderService_AddDish_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) UpdateDish(ctx context.Context, in *UpdateDishRequest, opts ...grpc.CallOption) (*UpdateDishResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateDishResponse)
	err := c.cc.Invoke(ctx, OrderService_UpdateDish_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) DeleteDish(ctx context.Context, in *DeleteDishRequest, opts ...grpc.CallOption) (*DeleteDishResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteDishResponse)
	err := c.cc.Invoke(ctx, OrderService_DeleteDish_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) ListDishes(ctx context.Context, in *ListDishesRequest, opts ...grpc.CallOption) (*ListDishesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListDishesResponse)
	err := c.cc.Invoke(ctx, OrderService_ListDishes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateOrderResponse)
	err := c.cc.Invoke(ctx, OrderService_CreateOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) UpdateOrderStatus(ctx context.Context, in *UpdateOrderStatusRequest, opts ...grpc.CallOption) (*UpdateOrderStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateOrderStatusResponse)
	err := c.cc.Invoke(ctx, OrderService_UpdateOrderStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) ListOfOrders(ctx context.Context, in *ListOfOrdersRequest, opts ...grpc.CallOption) (*ListOfOrdersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListOfOrdersResponse)
	err := c.cc.Invoke(ctx, OrderService_ListOfOrders_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetOrderByKitchenId(ctx context.Context, in *GetOrderByKitchenIdRequest, opts ...grpc.CallOption) (*GetOrderByKitchenIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetOrderByKitchenIdResponse)
	err := c.cc.Invoke(ctx, OrderService_GetOrderByKitchenId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) AddReview(ctx context.Context, in *AddReviewRequest, opts ...grpc.CallOption) (*AddReviewResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddReviewResponse)
	err := c.cc.Invoke(ctx, OrderService_AddReview_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) ListReviews(ctx context.Context, in *ListReviewsRequest, opts ...grpc.CallOption) (*ListReviewsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListReviewsResponse)
	err := c.cc.Invoke(ctx, OrderService_ListReviews_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) CreatePayment(ctx context.Context, in *CreatePaymentRequest, opts ...grpc.CallOption) (*CreatePaymentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreatePaymentResponse)
	err := c.cc.Invoke(ctx, OrderService_CreatePayment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetFullInfoAboutOrder(ctx context.Context, in *GetFullInfoAboutOrderRequest, opts ...grpc.CallOption) (*GetFullInfoAboutOrderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetFullInfoAboutOrderResponse)
	err := c.cc.Invoke(ctx, OrderService_GetFullInfoAboutOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetDishRecommendations(ctx context.Context, in *GetDishRecommendationsRequest, opts ...grpc.CallOption) (*GetDishRecommendationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDishRecommendationsResponse)
	err := c.cc.Invoke(ctx, OrderService_GetDishRecommendations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetKitchenStatistics(ctx context.Context, in *GetKitchenStatisticsRequest, opts ...grpc.CallOption) (*GetKitchenStatisticsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetKitchenStatisticsResponse)
	err := c.cc.Invoke(ctx, OrderService_GetKitchenStatistics_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetUserActivity(ctx context.Context, in *GetUserActivityRequest, opts ...grpc.CallOption) (*GetUserActivityResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserActivityResponse)
	err := c.cc.Invoke(ctx, OrderService_GetUserActivity_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) CreateKitchenWorkingHours(ctx context.Context, in *CreateKitchenWorkingHoursRequest, opts ...grpc.CallOption) (*CreateKitchenWorkingHoursResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateKitchenWorkingHoursResponse)
	err := c.cc.Invoke(ctx, OrderService_CreateKitchenWorkingHours_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) UpdateKitchenWorkingHours(ctx context.Context, in *UpdateKitchenWorkingHoursRequest, opts ...grpc.CallOption) (*UpdateKitchenWorkingHoursResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateKitchenWorkingHoursResponse)
	err := c.cc.Invoke(ctx, OrderService_UpdateKitchenWorkingHours_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) CreateDishNutritionInfo(ctx context.Context, in *CreateDishNutritionInfoRequest, opts ...grpc.CallOption) (*CreateDishNutritionInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateDishNutritionInfoResponse)
	err := c.cc.Invoke(ctx, OrderService_CreateDishNutritionInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) UpdateDishNutritionInfo(ctx context.Context, in *UpdateDishNutritionInfoRequest, opts ...grpc.CallOption) (*UpdateDishNutritionInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateDishNutritionInfoResponse)
	err := c.cc.Invoke(ctx, OrderService_UpdateDishNutritionInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServiceServer is the server API for OrderService service.
// All implementations must embed UnimplementedOrderServiceServer
// for forward compatibility
type OrderServiceServer interface {
	// Dish
	// 1 Done
	AddDish(context.Context, *AddDishRequest) (*AddDishResponse, error)
	// 2 Done
	UpdateDish(context.Context, *UpdateDishRequest) (*UpdateDishResponse, error)
	// 3 Done
	DeleteDish(context.Context, *DeleteDishRequest) (*DeleteDishResponse, error)
	// 4 Done
	ListDishes(context.Context, *ListDishesRequest) (*ListDishesResponse, error)
	// Order
	// 5 Done
	CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error)
	// 6 Done
	UpdateOrderStatus(context.Context, *UpdateOrderStatusRequest) (*UpdateOrderStatusResponse, error)
	// 7 Done
	ListOfOrders(context.Context, *ListOfOrdersRequest) (*ListOfOrdersResponse, error)
	// 8 Done
	GetOrderByKitchenId(context.Context, *GetOrderByKitchenIdRequest) (*GetOrderByKitchenIdResponse, error)
	// 9 Done
	AddReview(context.Context, *AddReviewRequest) (*AddReviewResponse, error)
	// 10 Done
	ListReviews(context.Context, *ListReviewsRequest) (*ListReviewsResponse, error)
	// 11
	CreatePayment(context.Context, *CreatePaymentRequest) (*CreatePaymentResponse, error)
	// 12
	GetFullInfoAboutOrder(context.Context, *GetFullInfoAboutOrderRequest) (*GetFullInfoAboutOrderResponse, error)
	// Qo'shimchalar
	// 1
	GetDishRecommendations(context.Context, *GetDishRecommendationsRequest) (*GetDishRecommendationsResponse, error)
	// 2
	GetKitchenStatistics(context.Context, *GetKitchenStatisticsRequest) (*GetKitchenStatisticsResponse, error)
	// 3
	GetUserActivity(context.Context, *GetUserActivityRequest) (*GetUserActivityResponse, error)
	// 4
	CreateKitchenWorkingHours(context.Context, *CreateKitchenWorkingHoursRequest) (*CreateKitchenWorkingHoursResponse, error)
	// 5
	UpdateKitchenWorkingHours(context.Context, *UpdateKitchenWorkingHoursRequest) (*UpdateKitchenWorkingHoursResponse, error)
	// 6
	CreateDishNutritionInfo(context.Context, *CreateDishNutritionInfoRequest) (*CreateDishNutritionInfoResponse, error)
	// 7
	UpdateDishNutritionInfo(context.Context, *UpdateDishNutritionInfoRequest) (*UpdateDishNutritionInfoResponse, error)
	mustEmbedUnimplementedOrderServiceServer()
}

// UnimplementedOrderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrderServiceServer struct {
}

func (UnimplementedOrderServiceServer) AddDish(context.Context, *AddDishRequest) (*AddDishResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDish not implemented")
}
func (UnimplementedOrderServiceServer) UpdateDish(context.Context, *UpdateDishRequest) (*UpdateDishResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDish not implemented")
}
func (UnimplementedOrderServiceServer) DeleteDish(context.Context, *DeleteDishRequest) (*DeleteDishResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDish not implemented")
}
func (UnimplementedOrderServiceServer) ListDishes(context.Context, *ListDishesRequest) (*ListDishesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDishes not implemented")
}
func (UnimplementedOrderServiceServer) CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedOrderServiceServer) UpdateOrderStatus(context.Context, *UpdateOrderStatusRequest) (*UpdateOrderStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrderStatus not implemented")
}
func (UnimplementedOrderServiceServer) ListOfOrders(context.Context, *ListOfOrdersRequest) (*ListOfOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOfOrders not implemented")
}
func (UnimplementedOrderServiceServer) GetOrderByKitchenId(context.Context, *GetOrderByKitchenIdRequest) (*GetOrderByKitchenIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderByKitchenId not implemented")
}
func (UnimplementedOrderServiceServer) AddReview(context.Context, *AddReviewRequest) (*AddReviewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddReview not implemented")
}
func (UnimplementedOrderServiceServer) ListReviews(context.Context, *ListReviewsRequest) (*ListReviewsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListReviews not implemented")
}
func (UnimplementedOrderServiceServer) CreatePayment(context.Context, *CreatePaymentRequest) (*CreatePaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePayment not implemented")
}
func (UnimplementedOrderServiceServer) GetFullInfoAboutOrder(context.Context, *GetFullInfoAboutOrderRequest) (*GetFullInfoAboutOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFullInfoAboutOrder not implemented")
}
func (UnimplementedOrderServiceServer) GetDishRecommendations(context.Context, *GetDishRecommendationsRequest) (*GetDishRecommendationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDishRecommendations not implemented")
}
func (UnimplementedOrderServiceServer) GetKitchenStatistics(context.Context, *GetKitchenStatisticsRequest) (*GetKitchenStatisticsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKitchenStatistics not implemented")
}
func (UnimplementedOrderServiceServer) GetUserActivity(context.Context, *GetUserActivityRequest) (*GetUserActivityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserActivity not implemented")
}
func (UnimplementedOrderServiceServer) CreateKitchenWorkingHours(context.Context, *CreateKitchenWorkingHoursRequest) (*CreateKitchenWorkingHoursResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateKitchenWorkingHours not implemented")
}
func (UnimplementedOrderServiceServer) UpdateKitchenWorkingHours(context.Context, *UpdateKitchenWorkingHoursRequest) (*UpdateKitchenWorkingHoursResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateKitchenWorkingHours not implemented")
}
func (UnimplementedOrderServiceServer) CreateDishNutritionInfo(context.Context, *CreateDishNutritionInfoRequest) (*CreateDishNutritionInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDishNutritionInfo not implemented")
}
func (UnimplementedOrderServiceServer) UpdateDishNutritionInfo(context.Context, *UpdateDishNutritionInfoRequest) (*UpdateDishNutritionInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDishNutritionInfo not implemented")
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

func _OrderService_AddDish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).AddDish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_AddDish_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).AddDish(ctx, req.(*AddDishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_UpdateDish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).UpdateDish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_UpdateDish_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).UpdateDish(ctx, req.(*UpdateDishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_DeleteDish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).DeleteDish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_DeleteDish_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).DeleteDish(ctx, req.(*DeleteDishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_ListDishes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDishesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).ListDishes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_ListDishes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).ListDishes(ctx, req.(*ListDishesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_CreateOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CreateOrder(ctx, req.(*CreateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_UpdateOrderStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrderStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).UpdateOrderStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_UpdateOrderStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).UpdateOrderStatus(ctx, req.(*UpdateOrderStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_ListOfOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOfOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).ListOfOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_ListOfOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).ListOfOrders(ctx, req.(*ListOfOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetOrderByKitchenId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderByKitchenIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetOrderByKitchenId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_GetOrderByKitchenId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetOrderByKitchenId(ctx, req.(*GetOrderByKitchenIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_AddReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).AddReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_AddReview_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).AddReview(ctx, req.(*AddReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_ListReviews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReviewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).ListReviews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_ListReviews_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).ListReviews(ctx, req.(*ListReviewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_CreatePayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CreatePayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_CreatePayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CreatePayment(ctx, req.(*CreatePaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetFullInfoAboutOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFullInfoAboutOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetFullInfoAboutOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_GetFullInfoAboutOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetFullInfoAboutOrder(ctx, req.(*GetFullInfoAboutOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetDishRecommendations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDishRecommendationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetDishRecommendations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_GetDishRecommendations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetDishRecommendations(ctx, req.(*GetDishRecommendationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetKitchenStatistics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKitchenStatisticsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetKitchenStatistics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_GetKitchenStatistics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetKitchenStatistics(ctx, req.(*GetKitchenStatisticsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetUserActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserActivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetUserActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_GetUserActivity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetUserActivity(ctx, req.(*GetUserActivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_CreateKitchenWorkingHours_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateKitchenWorkingHoursRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CreateKitchenWorkingHours(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_CreateKitchenWorkingHours_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CreateKitchenWorkingHours(ctx, req.(*CreateKitchenWorkingHoursRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_UpdateKitchenWorkingHours_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateKitchenWorkingHoursRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).UpdateKitchenWorkingHours(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_UpdateKitchenWorkingHours_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).UpdateKitchenWorkingHours(ctx, req.(*UpdateKitchenWorkingHoursRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_CreateDishNutritionInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDishNutritionInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CreateDishNutritionInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_CreateDishNutritionInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CreateDishNutritionInfo(ctx, req.(*CreateDishNutritionInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_UpdateDishNutritionInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDishNutritionInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).UpdateDishNutritionInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_UpdateDishNutritionInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).UpdateDishNutritionInfo(ctx, req.(*UpdateDishNutritionInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderService_ServiceDesc is the grpc.ServiceDesc for OrderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddDish",
			Handler:    _OrderService_AddDish_Handler,
		},
		{
			MethodName: "UpdateDish",
			Handler:    _OrderService_UpdateDish_Handler,
		},
		{
			MethodName: "DeleteDish",
			Handler:    _OrderService_DeleteDish_Handler,
		},
		{
			MethodName: "ListDishes",
			Handler:    _OrderService_ListDishes_Handler,
		},
		{
			MethodName: "CreateOrder",
			Handler:    _OrderService_CreateOrder_Handler,
		},
		{
			MethodName: "UpdateOrderStatus",
			Handler:    _OrderService_UpdateOrderStatus_Handler,
		},
		{
			MethodName: "ListOfOrders",
			Handler:    _OrderService_ListOfOrders_Handler,
		},
		{
			MethodName: "GetOrderByKitchenId",
			Handler:    _OrderService_GetOrderByKitchenId_Handler,
		},
		{
			MethodName: "AddReview",
			Handler:    _OrderService_AddReview_Handler,
		},
		{
			MethodName: "ListReviews",
			Handler:    _OrderService_ListReviews_Handler,
		},
		{
			MethodName: "CreatePayment",
			Handler:    _OrderService_CreatePayment_Handler,
		},
		{
			MethodName: "GetFullInfoAboutOrder",
			Handler:    _OrderService_GetFullInfoAboutOrder_Handler,
		},
		{
			MethodName: "GetDishRecommendations",
			Handler:    _OrderService_GetDishRecommendations_Handler,
		},
		{
			MethodName: "GetKitchenStatistics",
			Handler:    _OrderService_GetKitchenStatistics_Handler,
		},
		{
			MethodName: "GetUserActivity",
			Handler:    _OrderService_GetUserActivity_Handler,
		},
		{
			MethodName: "CreateKitchenWorkingHours",
			Handler:    _OrderService_CreateKitchenWorkingHours_Handler,
		},
		{
			MethodName: "UpdateKitchenWorkingHours",
			Handler:    _OrderService_UpdateKitchenWorkingHours_Handler,
		},
		{
			MethodName: "CreateDishNutritionInfo",
			Handler:    _OrderService_CreateDishNutritionInfo_Handler,
		},
		{
			MethodName: "UpdateDishNutritionInfo",
			Handler:    _OrderService_UpdateDishNutritionInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order-service/order-service.proto",
}
