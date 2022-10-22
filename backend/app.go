package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var name, _ = os.Hostname()
	result := add(1, 2)

	fmt.Fprintf(w, "<h1>This request was processed by host: %s</h1>\n", name)
	fmt.Fprintf(w, "<p>the result is: %d</p>\n", result)
}

func main() {
	postgresHost := os.Getenv("POSTGRES_HOST")
	postgresPort := os.Getenv("POSTGRES_PORT")
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassw := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		postgresHost, postgresPort, postgresUser, postgresPassw, dbName)
	_, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(os.Stdout, "Web Server started. Listening on 0.0.0.0:8080\n")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func add(a int, b int) int {
	return a + b
}
