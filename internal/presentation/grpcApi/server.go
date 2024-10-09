package grpcApi

import (
	"net"

	"github.com/abc-valera/netsly-golang/gen/pb"
	"github.com/abc-valera/netsly-golang/internal/application"
	"github.com/abc-valera/netsly-golang/internal/domain/global"
	"github.com/abc-valera/netsly-golang/internal/domain/service"
	"github.com/abc-valera/netsly-golang/internal/domain/util/coderr"
	"github.com/abc-valera/netsly-golang/internal/presentation/grpcApi/handler"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func RunServer(
	port string,
	_ service.Services,
	usecases application.Usecases,
) (
	serverStart func(),
	serverGracefulStop func(),
) {
	// Init handlers
	signHandler := handler.NewSignHandler(usecases.SignUsecase)

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
