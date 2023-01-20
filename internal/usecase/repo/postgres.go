package repo

import (
	"context"
	"fmt"
	"time"

	"gitlab.mapcard.pro/external-map-team/order-service/pkg/logger"
	"gitlab.mapcard.pro/external-map-team/order-service/pkg/metrics"
	"gorm.io/gorm"

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

func (r *OrderRepo) StoreOrder(ctx context.Context, req *entity.Order) (uint64, error) {
	beginTime := time.Now()

	defer func() {
		metrics.SetRequestTime("StoreOrder", float64(time.Since(beginTime).Milliseconds()))
		r.l.Infof("StoreOrder time %d", int(time.Since(beginTime).Milliseconds()))
	}()

	err := r.DB.Table("order").Transaction(func(tx *gorm.DB) error {
		return r.DB.Table("order").Create(&req).Error
	})

	if err != nil {
		return 0, fmt.Errorf("OrderRepo - StoreOrder - : %w", err)
	}

	return req.ID, nil
}
