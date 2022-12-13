package order

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgtype"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gitlab.mapcard.pro/external-map-team/api-proto/payment/api"
	orderpb "gitlab.mapcard.pro/external-map-team/order-service/app/proto"
)

// AddOrder adds a order to the directory.
func (d Directory) AddOrder(ctx context.Context, req *api.OrderRequest) (*api.Order, error) {
	fmt.Println("Start AddOrder..")
	orderRequest, err := json.Marshal(*req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error marshalling order request: %s", err.Error())
	}

	if req.MerchantOrderId == "" {
		return &api.Order{
			Success:    false,
			ErrCode:    "INVALID_FORMAT",
			ErrMessage: "Не указан уникальный номер заказа",
		}, nil
	}

	if req.Key == "" {
		return &api.Order{
			Success:    false,
			ErrCode:    "INVALID_FORMAT",
			ErrMessage: "Не указан уникальный индекс продавца",
		}, nil
	}

	var pgOrder Order
	pgOrder, err = d.querier.AddOrder(ctx, AddOrderParams{
		OrderRequest: json.RawMessage(orderRequest),
		Rrn:          "{}",
		OrderID:      req.MerchantOrderId,
		SellerID:     req.Key,
	})

	if err != nil {
		status.Errorf(codes.Internal, "unexpected error adding partner: %s", err.Error())
		return &api.Order{
			Success:    false,
			ErrCode:    "INTERNAL_ERROR",
			ErrMessage: err.Error(),
		}, nil
	}

	return orderPostgresToProto(pgOrder)
}

// ListOrders lists orders in the directory, subject to the request filters.
func (d Directory) ListOrders(req *api.ListOrdersRequest, srv orderpb.OrderService_ListOrdersServer) (retErr error) {
	q := d.sb.Select(
		"id",
		"create_time",
		"order_request",
		"rrn",
		"order_id",
		"seller_id",
	).From(
		"orders",
	)

	if req.GetCreatedSince() != nil {
		var pgTime pgtype.Timestamptz
		err := pgTime.Set(req.GetCreatedSince().AsTime())
		if err != nil {
			return status.Errorf(codes.InvalidArgument, "invalid timestamp: %s", err.Error())
		}
		q = q.Where(squirrel.Gt{
			"create_time": pgTime,
		})
	}

	if req.GetOlderThan() != nil {
		var pgInterval pgtype.Interval
		err := pgInterval.Set(req.GetOlderThan().AsDuration())
		if err != nil {
			return status.Errorf(codes.InvalidArgument, "invalid duration: %s", err.Error())
		}
		q = q.Where(
			squirrel.Expr(
				"CURRENT_TIMESTAMP - create_time > ?", pgInterval,
			),
		)
	}

	rows, retErr := q.QueryContext(srv.Context())
	if retErr != nil {
		return status.Error(codes.Internal, retErr.Error())
	}
	defer func() {
		cerr := rows.Close()
		if retErr == nil && cerr != nil {
			retErr = status.Error(codes.Internal, cerr.Error())
		}
	}()

	for rows.Next() {
		var pgOrder Order
		err := rows.Scan(
			&pgOrder.ID,
			&pgOrder.CreateTime,
			&pgOrder.OrderRequest,
			&pgOrder.Rrn,
			&pgOrder.OrderID,
			&pgOrder.SellerID,
		)
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}
		protoPartner, err := orderPostgresToProto(pgOrder)
		if err != nil {
			return err
		}
		err = srv.Send(protoPartner)
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}
	}

	retErr = rows.Err()
	if retErr != nil {
		return status.Error(codes.Internal, retErr.Error())
	}

	return nil
}

// DeleteOrder deletes the order, if found.
func (d Directory) DeleteOrder(ctx context.Context, req *api.SelectOrderRequest) (*api.Order, error) {
	pgOrder, err := d.querier.DeleteOrder(ctx, DeleteOrderParams{
		OrderID:  req.MerchantOrderId,
		SellerID: req.Key,
	})
	if err != nil {
		return &api.Order{
			Success:    false,
			ErrCode:    "NOT_FOUND",
			ErrMessage: "Заказ не найден",
		}, nil
	}
	return orderPostgresToProto(pgOrder)
}

// UpdateOrder updates order, if found.
func (d Directory) UpdateOrder(ctx context.Context, req *api.UpdateOrderRequest) (*api.Order, error) {
	pgOrder, err := d.querier.UpdateOrder(ctx, UpdateOrderParams{
		OrderID:  req.MerchantOrderId,
		SellerID: req.Key,
		Rrn:      req.Rrn,
	})
	if err != nil {
		return &api.Order{
			Success:    false,
			ErrCode:    "NOT_FOUND",
			ErrMessage: "Заказ не найден",
		}, nil
	}
	return orderPostgresToProto(pgOrder)
}

// GetOrder , if found.
func (d Directory) GetOrder(ctx context.Context, req *api.SelectOrderRequest) (*api.Order, error) {
	pgOrderRaw, err := d.querier.GetOrder(ctx, GetOrderParams{
		OrderID:  req.MerchantOrderId,
		SellerID: req.Key,
	})
	if err != nil {
		return &api.Order{
			Success:    false,
			ErrCode:    "NOT_FOUND",
			ErrMessage: "Заказ не найден",
		}, nil
	}

	return orderRawPostgresToProto(pgOrderRaw)
}
