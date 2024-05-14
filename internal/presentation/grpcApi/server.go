package grpcApi

import (
	"net"

	"github.com/abc-valera/netsly-api-golang/gen/pb"
	"github.com/abc-valera/netsly-api-golang/internal/application"
	"github.com/abc-valera/netsly-api-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/core/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain"
	"github.com/abc-valera/netsly-api-golang/internal/presentation/grpcApi/handler"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func RunServer(
	port string,
	staicPath string,

	services domain.Services,
	usecases application.UseCases,
) (
	serverStart func(),
	serverGracefulStop func(),
) {
	// Init handlers
	signHandler := handler.NewSignHandler(usecases.SignUseCase)

	// Register handlers
	server := grpc.NewServer()
	pb.RegisterSignServiceServer(server, signHandler)

	// ! Register reflection service on gRPC server (for development only)
	reflection.Register(server)

	lis := coderr.Must(net.Listen("tcp", port))

	return func() {
			global.Log().Info("grpcApi is running", "port", port)
			if err := server.Serve(lis); err != nil {
				coderr.Fatal("grpcApi server error: ", err)
			}
		}, func() {
			server.GracefulStop()
		}
}
