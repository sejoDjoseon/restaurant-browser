package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var name, _ = os.Hostname()

	fmt.Fprintf(w, "<h1>This request was processed by host: %s</h1>\n", name)
}

func (s *HttpServer) routes() {
	r := mux.NewRouter()
	r.HandleFunc("/restaurants", handler).Methods("GET")

	s.svr.Handler = r
}
