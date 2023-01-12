package usecase

import (
	"context"
	"time"

	"gitlab.mapcard.pro/external-map-team/order-service/internal/entity"
	"gitlab.mapcard.pro/external-map-team/order-service/pkg/logger"
	"gitlab.mapcard.pro/external-map-team/order-service/pkg/metrics"
)

type OrderUseCase struct {
	logger *logger.Logger
	repo   OrderRepo
}

func New(logger *logger.Logger, r OrderRepo) *OrderUseCase {
	return &OrderUseCase{
		logger: logger,
		repo:   r,
	}
}

func (u *OrderUseCase) AddOrder(ctx context.Context, order *entity.OrderCreator) (uint64, error) {
	beginTime := time.Now()

	defer func() {
		metrics.SetRequestTime("AddOrder", float64(time.Since(beginTime).Milliseconds()))
		u.logger.Infof("AddOrder time %d", int(time.Since(beginTime).Milliseconds()))
	}()

	id, err := u.repo.StoreOrder(ctx, order)

	if err != nil {
		return 0, err
	}

	return id, nil
}
