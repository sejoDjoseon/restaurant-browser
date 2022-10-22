package cmd

import (
	"context"
	"fmt"
	"os"
	"restaurant-browser/env"
	"restaurant-browser/internal/database"
	"restaurant-browser/server"

	_ "github.com/lib/pq"
)

func Start() {
	ctx := context.Background()

	envApp := env.GetEnv()

	db := database.NewPostgresSQL(ctx, envApp)
	err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.CloseDB()

	server := server.NewHttpServer(ctx, envApp, db)
	err = server.Run()
	if err != nil {
		fmt.Fprintf(os.Stdout, "Server stopped: %s\n", err.Error())
	}
}
