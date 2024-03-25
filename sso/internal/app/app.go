package app

import (
	grpcapp "github.com/17HIERARCH70/messageService/sso/internal/app/grpc"
	"github.com/17HIERARCH70/messageService/sso/internal/config"
	"github.com/17HIERARCH70/messageService/sso/internal/services/auth"
	"github.com/17HIERARCH70/messageService/sso/internal/storage/postgresql"
	"log/slog"
	"time"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	tokenTTL time.Duration,
	dbConfig *config.PostgresSQLConfig,
) *App {
	db, err := postgresql.Connect(dbConfig)
	if err != nil {
		panic(err)
	}

	storage := postgresql.NewStorage(db)

	authService := auth.New(log, storage, storage, storage, tokenTTL)

	grpcApp := grpcapp.New(log, authService, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
