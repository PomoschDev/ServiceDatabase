package app

import (
	grpcapp "DababaseService/iternal/app/grpc"
	"log/slog"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(log *slog.Logger, grpcPort int) *App {
	grpcApp := grpcapp.New(log, grpcPort)

	return &App{GRPCSrv: grpcApp}
}
