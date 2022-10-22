package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"restaurant-browser/env"
	"restaurant-browser/internal/database"
)

type HttpServer struct {
	svr *http.Server
	ctx context.Context
	db  database.Database
	ec  env.EnvApp
}

func NewHttpServer(ctx context.Context, ec env.EnvApp, db database.Database) Server {
	httpServer := &http.Server{
		Addr: (fmt.Sprintf(":%s", ec.SERVER_PORT)),
	}

	server := &HttpServer{
		svr: httpServer,
		ctx: ctx,
		db:  db,
		ec:  ec,
	}

	server.routes()

	return server
}

func (s *HttpServer) Run() error {
	fmt.Fprintf(os.Stdout, "Web Server started. Listening on 0.0.0.0:%s\n", s.ec.SERVER_PORT)
	return s.svr.ListenAndServe()
}
