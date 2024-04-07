package app

import (
	grpcapp "github.com/17HIERARCH70/gRPCserviceExample/sso/internal/app/grpc"
	"github.com/17HIERARCH70/gRPCserviceExample/sso/internal/services/auth"
	"github.com/17HIERARCH70/gRPCserviceExample/sso/internal/storage/sqlite"
	"log/slog"
	"time"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration,
) *App {
	storage, err := sqlite.New(storagePath)
	if err != nil {
		panic(err)
	}

	authService := auth.New(log, storage, storage, storage, tokenTTL)

	grpcApp := grpcapp.New(log, authService, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
