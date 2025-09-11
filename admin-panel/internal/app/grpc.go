package app

import (
	"fmt"
	"net"
	s "admin-panel/internal/service"

	pb "github.com/xadichamakhkamova/HospitalContracts/genproto/adminpb"

	"google.golang.org/grpc"
)

type App struct {
	gRPCServer *grpc.Server
	port       int
}

func New(srv s.AdminService, port int) *App {
	grpcServer := grpc.NewServer()
	pb.RegisterAdminServiceServer(grpcServer, &srv)
	return &App{
		gRPCServer: grpcServer,
		port:       port,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	addr := fmt.Sprintf(":%d", a.port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	if err := a.gRPCServer.Serve(listener); err != nil {
		return err
	}
	return nil
}


func (a *App) Stop() {
	a.gRPCServer.GracefulStop()
}