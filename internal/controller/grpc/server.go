package grpc

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"time"

	"gitlab.mapcard.pro/external-map-team/api-proto/payment/api"
	"gitlab.mapcard.pro/external-map-team/order-service/internal/entity"
	"gitlab.mapcard.pro/external-map-team/order-service/internal/usecase"
	"gitlab.mapcard.pro/external-map-team/order-service/pkg/logger"
	"gitlab.mapcard.pro/external-map-team/order-service/pkg/metrics"
	"gitlab.mapcard.pro/external-map-team/order-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type Server struct {
	proto.UnimplementedOrderServiceServer
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

	proto.RegisterOrderServiceServer(grpcServer, s)

	grpc_health_v1.RegisterHealthServer(grpcServer, health.NewServer())

	s.logger.Infof("Start serve GRPC at :%s", port)

	go grpcServer.Serve(listener)

	return nil
}

func (s *Server) AddOrder(ctx context.Context, req *api.OrderRequest) (*api.Order, error) {
	beginTime := time.Now()

	defer func() {
		metrics.SetRequestTime("AddOrder GRPC", float64(time.Since(beginTime).Milliseconds()))
		s.logger.Infof("AddOrder GRPC time %d", int(time.Since(beginTime).Milliseconds()))
	}()

	s.logger.Infof("got grpc query AddOrder")

	itemsBytes, err := json.Marshal(req.Items)
	if err != nil {
		return nil, err
	}

	buyerBytes, err := json.Marshal(req.Buyer)
	if err != nil {
		return nil, err
	}

	userDataBytes, err := json.Marshal(req.UserData)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	var items []*api.Item
	err = json.Unmarshal([]byte(order.Items), &items)

	if err != nil {
		return nil, err
	}

	var buyer api.Buyer
	err = json.Unmarshal([]byte(order.Buyer), &buyer)

	if err != nil {
		return nil, err
	}

	response := api.Order{
		MerchantOrderId: order.MerchantOrderId,
		Currency:        order.Currency,
		Amount:          uint32(order.Amount),
		PaymentType:     api.PayType(order.PaymentType),
		Items:           items,
		Buyer:           &buyer,
		OrderId:         newId,
	}

	return &response, nil
}
