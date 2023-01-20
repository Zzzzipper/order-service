package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"gitlab.mapcard.pro/external-map-team/order-service/config"
	"gitlab.mapcard.pro/external-map-team/order-service/internal/controller/grpc"
	"gitlab.mapcard.pro/external-map-team/order-service/internal/usecase"
	"gitlab.mapcard.pro/external-map-team/order-service/internal/usecase/repo"
	"gitlab.mapcard.pro/external-map-team/order-service/pkg/logger"
	"gitlab.mapcard.pro/external-map-team/order-service/pkg/postgres"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	pg, err := postgres.New(cfg.Postgres)

	if err != nil {
		l.Fatalf(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	orderUseCase := usecase.New(
		l,
		repo.New(pg, l),
	)

	grpcServer := grpc.NewServer(l, orderUseCase)

	err = grpcServer.Start(cfg.GRPC.Port)

	if err != nil {
		l.Fatalf(fmt.Errorf("app - NewApp - grpcServer start error: %w", err))
	}

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Infof("app - Run - signal: " + s.String())
	}

}
