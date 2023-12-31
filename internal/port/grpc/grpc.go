package grpc

import (
	"net"

	"github.com/abc-valera/flugo-api-golang/gen/pb"
	"github.com/abc-valera/flugo-api-golang/internal/core/application"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/service"
	"github.com/abc-valera/flugo-api-golang/internal/port/grpc/handler"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func RunServer(
	port string,
	repos repository.Repositories,
	services service.Services,
	usecases application.UseCases,
) error {
	// Init handlers
	signHandler := handler.NewSignHandler(repos.UserRepo, usecases.SignUseCase)

	// Register handlers
	server := grpc.NewServer()
	pb.RegisterSignServiceServer(server, signHandler)

	// ! Register reflection service on gRPC server (for development only)
	reflection.Register(server)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	service.Log.Info("Starting gRPC server on " + port)
	return server.Serve(lis)
}
