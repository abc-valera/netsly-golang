package grpcapi

import (
	"net"

	"github.com/abc-valera/netsly-api-golang/gen/pb"
	"github.com/abc-valera/netsly-api-golang/internal/application"
	"github.com/abc-valera/netsly-api-golang/internal/domain"
	"github.com/abc-valera/netsly-api-golang/internal/domain/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/presentation/grpc-api/handler"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func RunServer(
	presentation string,
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

	lis := coderr.MustWithVal(net.Listen("tcp", presentation))

	return func() {
			global.Log().Info("grpc-api is running", "presentation", presentation)
			if err := server.Serve(lis); err != nil {
				global.Log().Fatal("grpc-api server error: ", err)
			}
		}, func() {
			server.GracefulStop()
		}
}