package grpcApi

import (
	"github.com/abc-valera/netsly-golang/gen/pb"
	"github.com/abc-valera/netsly-golang/internal/application"
	"github.com/abc-valera/netsly-golang/internal/domain/entity"
	"github.com/abc-valera/netsly-golang/internal/domain/global"
	"github.com/abc-valera/netsly-golang/internal/domain/service"
	"github.com/abc-valera/netsly-golang/internal/presentation/grpcApi/handler"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewServer(
	_ entity.Entities,
	_ service.Services,
	usecases application.Usecases,
) *grpc.Server {
	// Init handlers
	signHandler := handler.NewSignHandler(usecases.SignUsecase)

	// Register handlers
	server := grpc.NewServer()
	pb.RegisterSignServiceServer(server, signHandler)

	// Register reflection service on gRPC server (for development only)
	if !global.IsProduction() {
		reflection.Register(server)
	}

	return server
}
