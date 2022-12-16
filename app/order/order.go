package order

import (
	"context"
	"encoding/json"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgtype"
	"google.golang.org/grpc/codes"

	"gitlab.mapcard.pro/external-map-team/api-proto/payment/api"
	orderpb "gitlab.mapcard.pro/external-map-team/order-service/app/proto"
)

// AddOrder adds a order to the directory.
func (d Directory) AddOrder(ctx context.Context, req *api.OrderRequest) (*api.Order, error) {
	Log().Line("Start AddOrder..")

	orderRequest, err := json.Marshal(*req)
	if err != nil {
		return nil, Log().StatusErrorf(codes.Internal, "Error marshalling order request: %s", err.Error())
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
		Log().StatusErrorf(codes.Internal, "Unexpected error adding partner: %s", err.Error())
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
			return Log().StatusErrorf(codes.InvalidArgument, "Invalid timestamp: %s", err.Error())
		}
		q = q.Where(squirrel.Gt{
			"create_time": pgTime,
		})
	}

	if req.GetOlderThan() != nil {
		var pgInterval pgtype.Interval
		err := pgInterval.Set(req.GetOlderThan().AsDuration())
		if err != nil {
			return Log().StatusErrorf(codes.InvalidArgument, "Invalid duration: %s", err.Error())
		}
		q = q.Where(
			squirrel.Expr(
				"CURRENT_TIMESTAMP - create_time > ?", pgInterval,
			),
		)
	}

	rows, retErr := q.QueryContext(srv.Context())
	if retErr != nil {
		return Log().StatusError(codes.Internal, retErr.Error())
	}
	defer func() {
		cerr := rows.Close()
		if retErr == nil && cerr != nil {
			retErr = Log().StatusError(codes.Internal, cerr.Error())
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
			return Log().StatusError(codes.Internal, err.Error())
		}
		protoPartner, err := orderPostgresToProto(pgOrder)
		if err != nil {
			return err
		}
		err = srv.Send(protoPartner)
		if err != nil {
			return Log().StatusError(codes.Internal, err.Error())
		}
	}

	retErr = rows.Err()
	if retErr != nil {
		return Log().StatusError(codes.Internal, retErr.Error())
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
