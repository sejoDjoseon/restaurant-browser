package server

import (
	app_restaurant "restaurant-browser/core/application"

	"github.com/gorilla/mux"
)

func (s *HttpServer) routes() {
	rtsService := app_restaurant.NewRestaurantHTTPService(s.ctx, s.db.GetConnection())
	r := mux.NewRouter()
	r.HandleFunc("/restaurants", rtsService.Handler).Methods("GET")

	s.svr.Handler = r
}
