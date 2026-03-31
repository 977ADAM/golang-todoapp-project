package corehttpserver

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	corelogger "github.com/977ADAM/golang-todoapp-project/internal/core/logger"
	corehttpmiddleware "github.com/977ADAM/golang-todoapp-project/internal/core/transport/http/middleware"
	"go.uber.org/zap"
)



type HTTPServer struct {
	mux *http.ServeMux
	config Config
	log *corelogger.Logger

	middleware []corehttpmiddleware.Middleware
}

func NewHTTPServer(
	config Config,
	log *corelogger.Logger,
	middleware ...corehttpmiddleware.Middleware,
) *HTTPServer {
	return &HTTPServer {
		mux: http.NewServeMux(),
		config: config,
		log: log,
		middleware: middleware,
	}
}

func (h *HTTPServer) RegisterAPIRoutes(routers ...*APIVersionRouter) {
	for _, router := range routers {
		prefix := "/api/" + string(router.apiVersion)

		h.mux.Handle(
			prefix+"/",
			http.StripPrefix(prefix, router),
		)
	}
}

func (h *HTTPServer) Run(ctx context.Context) error {
	mux := corehttpmiddleware.ChainMiddleware(h.mux, h.middleware...)

	server := http.Server {
		Addr: h.config.Addr,
		Handler: mux,
	}

	ch := make(chan error, 1)


	go func() {
		defer close(ch)

		h.log.Warn("start HTTP server", zap.String("addr", h.config.Addr))

		err := server.ListenAndServe()

		if !errors.Is(err, http.ErrServerClosed) {
			ch <- err
		}
	}()

	select {
	case err := <-ch:
		if err != nil {
			return fmt.Errorf("listen and server HTTP: %w", err)
		}
	case <-ctx.Done():
		h.log.Warn("shutdown HTTP server...")

		shutdownCtx, cancel := context.WithTimeout(
			context.Background(),
			h.config.ShutdownTimeout,
		)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			_ = server.Close()

			return fmt.Errorf("shutdown HTTP server: %w", err)
		}
		
		h.log.Warn("HTTP server stopped")
	}

	return nil
}