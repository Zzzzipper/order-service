package grpc

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"time"

	order_api "gitlab.mapcard.pro/external-map-team/api-proto/order/api"
	payment_api "gitlab.mapcard.pro/external-map-team/api-proto/payment/api"
	"gitlab.mapcard.pro/external-map-team/order-service/internal/entity"
	"gitlab.mapcard.pro/external-map-team/order-service/internal/usecase"
	"gitlab.mapcard.pro/external-map-team/order-service/pkg/logger"
	"gitlab.mapcard.pro/external-map-team/order-service/pkg/metrics"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type Server struct {
	order_api.UnimplementedOrderServiceServer
	logger *logger.Logger
	order  usecase.Order
}

func NewServer(logger *logger.Logger, order usecase.Order) *Server {
	return &Server{
		logger: logger,
		order:  order,
	}
}

func (s *Server) Start(port string) error {
	addr := fmt.Sprintf(":%s", port)

	listener, err := net.Listen("tcp", addr)

	if err != nil {
		s.logger.Errorf(err)
		return err
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	order_api.RegisterOrderServiceServer(grpcServer, s)

	grpc_health_v1.RegisterHealthServer(grpcServer, health.NewServer())

	s.logger.Infof("Start serve GRPC at :%s", port)

	go grpcServer.Serve(listener)

	return nil
}

func (s *Server) AddOrder(ctx context.Context, req *order_api.AddOrderRequest) (*order_api.AddOrderResponse, error) {
	beginTime := time.Now()

	defer func() {
		metrics.SetRequestTime("AddOrder GRPC", float64(time.Since(beginTime).Milliseconds()))
		s.logger.Infof("AddOrder GRPC time %d", int(time.Since(beginTime).Milliseconds()))
	}()

	s.logger.Infof("got grpc query AddOrder")

	itemsBytes, err := json.Marshal(req.Items)
	if err != nil {
		s.logger.Errorf("AddOrder: json.Marshal(req.Items) - %w", err)
		return &order_api.AddOrderResponse{
			Status:     false,
			ErrCode:    "INTERNAL_ERROR",
			ErrMessage: err.Error(),
		}, nil
	}

	buyerBytes, err := json.Marshal(req.Buyer)
	if err != nil {
		s.logger.Errorf("AddOrder: json.Marshal(req.Buyer) - %w", err)
		return &order_api.AddOrderResponse{
			Status:     false,
			ErrCode:    "INTERNAL_ERROR",
			ErrMessage: err.Error(),
		}, nil
	}

	userDataBytes, err := json.Marshal(req.UserData)
	if err != nil {
		s.logger.Errorf("AddOrder: json.Marshal(req.UserData) - %w", err)
		return &order_api.AddOrderResponse{
			Status:     false,
			ErrCode:    "INTERNAL_ERROR",
			ErrMessage: err.Error(),
		}, nil
	}

	order := &entity.Order{
		MerchantOrderId: req.MerchantOrderId,
		Currency:        req.Currency,
		PaymentType:     int32(req.PaymentType),
		Amount:          int64(req.Amount),
		Lifetime:        req.Lifetime.AsDuration(),
		Items:           string(itemsBytes),
		Buyer:           string(buyerBytes),
		UserData:        string(userDataBytes),
		ShopId:          int64(req.ShopId),
		MerchantId:      int64(req.MerchantId),
		PartnerId:       int64(req.PartnerId),
	}

	newId, err := s.order.AddOrder(ctx, order)

	if err != nil {
		s.logger.Errorf("AddOrder: s.order.AddOrder(ctx, order) - %w", err)
		return &order_api.AddOrderResponse{
			Status:     false,
			ErrCode:    "INTERNAL_ERROR",
			ErrMessage: err.Error(),
		}, nil
	}

	var items []*payment_api.Item
	err = json.Unmarshal([]byte(order.Items), &items)

	if err != nil {
		s.logger.Errorf("AddOrder: json.Unmarshal([]byte(order.Items), &items) - %w", err)
		return &order_api.AddOrderResponse{
			Status:     false,
			ErrCode:    "INTERNAL_ERROR",
			ErrMessage: err.Error(),
		}, nil
	}

	var buyer payment_api.Buyer
	err = json.Unmarshal([]byte(order.Buyer), &buyer)

	if err != nil {
		s.logger.Errorf("AddOrder: json.Unmarshal([]byte(order.Buyer), &buyer) - %w", err)
		return &order_api.AddOrderResponse{
			Status:     false,
			ErrCode:    "INTERNAL_ERROR",
			ErrMessage: err.Error(),
		}, nil
	}

	response := order_api.AddOrderResponse{
		Status:  true,
		ErrCode: "OK",
		Order: &order_api.Order{
			MerchantOrderId: order.MerchantOrderId,
			Currency:        order.Currency,
			Amount:          uint32(order.Amount),
			PaymentType:     order_api.PayType(order.PaymentType),
			Items:           items,
			Buyer:           &buyer,
			OrderId:         newId,
		},
	}

	return &response, nil
}
