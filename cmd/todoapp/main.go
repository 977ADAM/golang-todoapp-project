package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	corelogger "github.com/977ADAM/golang-todoapp-project/internal/core/logger"
	corepgxpool "github.com/977ADAM/golang-todoapp-project/internal/core/repository/postgres/pool/pgx"
	corehttpmiddleware "github.com/977ADAM/golang-todoapp-project/internal/core/transport/http/middleware"
	corehttpserver "github.com/977ADAM/golang-todoapp-project/internal/core/transport/http/server"
	userspostgresrepository "github.com/977ADAM/golang-todoapp-project/internal/features/users/repository/postgres"
	usersservice "github.com/977ADAM/golang-todoapp-project/internal/features/users/service"
	userstransporthttp "github.com/977ADAM/golang-todoapp-project/internal/features/users/transport/http"
	"go.uber.org/zap"
)

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM,
	)
	defer cancel()

	logger, err := corelogger.NewLogger(corelogger.NewConfigMust())
	if err != nil {
		fmt.Println("failed to init application logger:", err)
		os.Exit(1)
	}
	defer logger.Close()

	logger.Debug("initializing postgres connection pool")
	pool, err := corepgxpool.NewPool(
		ctx,
		corepgxpool.NewConfigMust(),
	)
	if err != nil {
		logger.Fatal("failed to init postgres connection pool", zap.Error(err))
	}
	defer pool.Close()

	logger.Debug("initializing feature", zap.String("feature", "users"))
	usersRepository := userspostgresrepository.NewUsersRepository(pool)
	usersService := usersservice.NewUsersService(usersRepository)

	usersTransportHTTP := userstransporthttp.NewUsersHTTPHandler(usersService)

	logger.Debug("initializing HTTP server")
	httpServer := corehttpserver.NewHTTPServer(
		corehttpserver.NewConfigMust(),
		logger,
		corehttpmiddleware.RequestID(),
		corehttpmiddleware.Logger(logger),
		corehttpmiddleware.Trace(),
		corehttpmiddleware.Panic(),
	)

	apiVersionRouterV1 := corehttpserver.NewAPIVersionRouter(corehttpserver.ApiVersion1)
	apiVersionRouterV1.RegisterRoutes(usersTransportHTTP.Routes()...)

	// apiVersionRouterV2 := corehttpserver.NewAPIVersionRouter(
	// 	corehttpserver.ApiVersion2,
	// 	corehttpmiddleware.Dummy("API v2 middleware"),
	// )
	// apiVersionRouterV2.RegisterRoutes(usersTransportHTTP.Routes()...)

	httpServer.RegisterAPIRoutes(
		apiVersionRouterV1,
		// apiVersionRouterV2,
	)

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server tun error", zap.Error(err))
	}

}
