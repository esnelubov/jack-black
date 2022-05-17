package server_http

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"os"
)

func RunHTTPServer(createHandler func(router chi.Router) http.Handler) {
	port := os.Getenv("BJ_PORT")
	if port == "" {
		port = "3434"
	}
	RunHTTPServerOnAddr(":"+port, createHandler)
}

func RunHTTPServerOnAddr(addr string, createHandler func(router chi.Router) http.Handler) {
	apiRouter := chi.NewRouter()
	setMiddlewares(apiRouter)

	rootRouter := chi.NewRouter()
	// we are mounting all APIs under /api path
	rootRouter.Mount("/api", createHandler(apiRouter))

	if err := http.ListenAndServe(addr, rootRouter); err != nil {
		panic(err)
	}
}

func setMiddlewares(router *chi.Mux) {
	// TODO Add middleware to check request according to swagger
	// TODO Add middleware to check authorization
}
