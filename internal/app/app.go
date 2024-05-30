package app

import (
	grpcapp "diplomServer/internal/app/grpc"
	"diplomServer/internal/grpc/auth"
	"diplomServer/internal/grpc/machine_department"
	"diplomServer/internal/grpc/radar"
	"diplomServer/internal/grpc/user"
	"diplomServer/internal/models"
	"diplomServer/internal/repository/postgreSQL"
	"go.uber.org/zap"
)

type App struct {
	GRPSrv *grpcapp.App
}

func New(log *zap.SugaredLogger, cfg *models.Config) *App {
	storage, err := postgreSQL.New(cfg)
	if err != nil {
		panic(err)
	}

	authService := auth.New(log, storage, cfg.TokenTTl)
	machineService := machine_department.New(log, storage, cfg.TokenTTl)
	userService := user.New(log, storage, cfg.TokenTTl)
	radarService := radar.New(log)

	grpcApp := grpcapp.New(log, authService, machineService, userService, radarService, cfg.ServicePort)

	return &App{
		GRPSrv: grpcApp,
	}
}
