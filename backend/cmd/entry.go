package cmd

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"restaurant-browser/env"

	_ "github.com/lib/pq"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var name, _ = os.Hostname()

	fmt.Fprintf(w, "<h1>This request was processed by host: %s</h1>\n", name)
}

func Start() {
	envApp := env.GetEnv()

	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		envApp.DB_HOST,
		envApp.DB_PORT,
		envApp.DB_USER,
		envApp.DB_PASSWORD,
		envApp.DB_NAME,
	)
	_, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(os.Stdout, "Web Server started. Listening on 0.0.0.0:%s\n", envApp.SERVER_PORT)
	http.HandleFunc("/", handler)
	http.ListenAndServe(fmt.Sprintf(":%s", envApp.SERVER_PORT), nil)
}
