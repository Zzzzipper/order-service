package repo

import (
	"context"
	"fmt"
	"time"

	"gitlab.mapcard.pro/external-map-team/order-service/pkg/logger"
	"gitlab.mapcard.pro/external-map-team/order-service/pkg/metrics"

	"gitlab.mapcard.pro/external-map-team/order-service/internal/entity"
	"gitlab.mapcard.pro/external-map-team/order-service/pkg/postgres"
)

type OrderRepo struct {
	*postgres.Postgres
	l *logger.Logger
}

func New(pg *postgres.Postgres, l *logger.Logger) *OrderRepo {
	return &OrderRepo{pg, l}
}

func (r *OrderRepo) StoreOrder(ctx context.Context, req *entity.OrderCreator) error {
	beginTime := time.Now()

	defer func() {
		metrics.SetRequestTime("StoreOrder", float64(time.Since(beginTime).Milliseconds()))
		r.l.Infof("StoreOrder time %d", int(time.Since(beginTime).Milliseconds()))
	}()

	result := r.DB.Table("order").Create(&req)

	r.DB.Table("order").Save(req)

	if result.Error != nil {
		return fmt.Errorf("OrderRepo - StoreOrder - error: %w", result.Error)
	}

	return nil
}
