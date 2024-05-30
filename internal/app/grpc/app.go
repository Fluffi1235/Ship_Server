package grpcapp

import (
	authgRPC "diplomServer/internal/grpc/auth"
	"diplomServer/internal/grpc/machine_department"
	"diplomServer/internal/grpc/radar"
	"diplomServer/internal/grpc/user"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

type App struct {
	log        *zap.SugaredLogger
	gRPCServer *grpc.Server
	port       int
}

func New(log *zap.SugaredLogger, authService *authgRPC.AuthImpl, machineService *machine_department.MachineImpl,
	userService *user.UserImpl, radarService *radar.RadarImpl, port int) *App {

	gRPCServer := grpc.NewServer()

	authgRPC.Register(gRPCServer, authService)
	machine_department.Register(gRPCServer, machineService)
	user.Register(gRPCServer, userService)
	radar.Register(gRPCServer, radarService)

	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

// Run runs gRPC server.
func (a *App) Run() error {
	const op = "grpcapp.Run"

	l, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", a.port))
	if err != nil {
		return errors.New(fmt.Sprintf("%s: %w", op, err.Error()))
	}

	a.log.Info(fmt.Sprintf("grpc server is running addr %s", l.Addr().String()))

	if err != nil {
		return errors.New(fmt.Sprintf("%s: %v", op, err))
	}

	err = a.gRPCServer.Serve(l)
	if err != nil {
		return errors.New(fmt.Sprintf("%s: %v", op, err))
	}

	return nil
}

// Stop stops gRPC server.
func (a *App) Stop() {
	a.log.Info("stopping gRPC server")

	a.gRPCServer.GracefulStop()
}
