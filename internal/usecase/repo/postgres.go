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

func (r *OrderRepo) StoreOrder(ctx context.Context, req *entity.OrderCreator) (uint64, error) {
	beginTime := time.Now()

	defer func() {
		metrics.SetRequestTime("StoreOrder", float64(time.Since(beginTime).Milliseconds()))
		r.l.Infof("StoreOrder time %d", int(time.Since(beginTime).Milliseconds()))
	}()

	result := r.DB.Table("order").Create(&req)

	idCatcher := entity.OrderCreator{}
	r.DB.Table("order").Save(&idCatcher)

	if result.Error != nil {
		return 0, fmt.Errorf("OrderRepo - StoreOrder - error: %w", result.Error)
	}

	return idCatcher.ID, nil
}
