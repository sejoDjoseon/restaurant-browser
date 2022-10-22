package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"restaurant-browser/env"
	"restaurant-browser/internal/database"

	_ "github.com/lib/pq"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var name, _ = os.Hostname()

	fmt.Fprintf(w, "<h1>This request was processed by host: %s</h1>\n", name)
}

func Start() {
	ctx := context.Background()

	envApp := env.GetEnv()

	db := database.NewPostgresSQL(ctx, envApp)
	err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.CloseDB()

	fmt.Fprintf(os.Stdout, "Web Server started. Listening on 0.0.0.0:%s\n", envApp.SERVER_PORT)
	http.HandleFunc("/", handler)
	http.ListenAndServe(fmt.Sprintf(":%s", envApp.SERVER_PORT), nil)
}
