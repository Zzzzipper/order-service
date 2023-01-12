package usecase

import (
	"context"

	"gitlab.mapcard.pro/external-map-team/order-service/internal/entity"
)

type (
	Order interface {
		AddOrder(ctx context.Context, req *entity.Order) (uint64, error)
	}

	OrderRepo interface {
		StoreOrder(ctx context.Context, req *entity.Order) (uint64, error)
	}
)
