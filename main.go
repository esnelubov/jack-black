package main

import (
	"context"
	"github.com/go-chi/chi/v5"
	"jackBlack/internal/app"
	"jackBlack/internal/common/server/server_http"
	"jackBlack/internal/ports"
	"jackBlack/internal/ports/http_api"
	"net/http"
)

func main() {
	ctx := context.Background()

	application := app.NewApplication(ctx)

	server_http.RunHTTPServer(func(router chi.Router) http.Handler {
		return http_api.HandlerFromMux(
			ports.NewHttpServer(application),
			router,
		)
	})
}
