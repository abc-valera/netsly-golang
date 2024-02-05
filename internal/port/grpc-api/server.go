package grpcapi

import (
	"net"

	"github.com/abc-valera/netsly-api-golang/gen/pb"
	"github.com/abc-valera/netsly-api-golang/internal/application"
	"github.com/abc-valera/netsly-api-golang/internal/domain"
	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/port/grpc-api/handler"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// TODO: remade to the new standard
func RunServer(
	port string,
	services domain.Services,
	usecases application.UseCases,
) error {
	// Init handlers
	signHandler := handler.NewSignHandler(usecases.SignUseCase)

	// Register handlers
	server := grpc.NewServer()
	pb.RegisterSignServiceServer(server, signHandler)

	// ! Register reflection service on gRPC server (for development only)
	reflection.Register(server)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	global.Log.Info("Starting gRPC server on " + port)
	return server.Serve(lis)
}
