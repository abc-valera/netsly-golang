package presentation

import (
	"context"
	"errors"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/abc-valera/netsly-golang/internal/application"
	"github.com/abc-valera/netsly-golang/internal/domain/entity"
	"github.com/abc-valera/netsly-golang/internal/domain/global"
	"github.com/abc-valera/netsly-golang/internal/domain/service"
	"github.com/abc-valera/netsly-golang/internal/domain/util/coderr"
	"github.com/abc-valera/netsly-golang/internal/domain/util/env"
	"github.com/abc-valera/netsly-golang/internal/presentation/grpcApi"
	"github.com/abc-valera/netsly-golang/internal/presentation/jsonApi"
	"github.com/abc-valera/netsly-golang/internal/presentation/webApp"
)

func StartServer(
	services service.Services,
	entities entity.Entities,
	usecases application.Usecases,
) {
	// Initialize http server
	webAppHandler := webApp.NewServer(
		env.Load("WEB_APP_TEMPLATE_PATH"),
		services,
		entities,
		usecases,
	)

	jsonApiHandler := jsonApi.NewServer(
		env.Load("JWT_SIGN_KEY"),
		env.LoadDuration("ACCESS_TOKEN_DURATION"),
		env.LoadDuration("REFRESH_TOKEN_DURATION"),
		entities,
		services,
		usecases,
	)

	httpServer := &http.Server{
		Addr: env.Load("HTTP_PORT"),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			switch r.Host {
			case global.SubdomainWebApp() + global.DomainName():
				webAppHandler.ServeHTTP(w, r)
			case global.SubdomainJsonApi() + global.DomainName():
				jsonApiHandler.ServeHTTP(w, r)
			default:
				http.Error(w, "Forbidden subdomain", http.StatusForbidden)
			}
		}),
	}

	// Initialize grpc server
	grpcServer := grpcApi.NewServer(
		entities,
		services,
		usecases,
	)

	// Start the servers inside goroutines
	go func() {
		global.Log().Info("http server is running", "port", env.Load("HTTP_PORT"))
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			coderr.Fatal("jsonApi server error", err)
		}
	}()

	go func() {
		global.Log().Info("grpc server is running", "port", env.Load("GRPC_PORT"))
		if err := grpcServer.Serve(coderr.Must(net.Listen("tcp", env.Load("GRPC_PORT")))); err != nil {
			coderr.Fatal("grpc server error", err)
		}
	}()

	// Stop program execution until receiving an interrupt signal
	gracefulShutdown := make(chan os.Signal, 1)
	signal.Notify(gracefulShutdown, os.Interrupt)
	<-gracefulShutdown

	// Shutdown the http server
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		global.Log().Error("http server shutdown error", "error", err)
	}

	// Shutdown the grpc server
	grpcServer.GracefulStop()
}
