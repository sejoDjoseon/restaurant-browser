package app_restaurant

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	restaurant "restaurant-browser/core/infrastructure"
)

type RestaurantHTTPService struct {
	gtw restaurant.RestaurantGateway
}

func NewRestaurantHTTPService(ctx context.Context, db *sql.DB) *RestaurantHTTPService {
	return &RestaurantHTTPService{restaurant.NewRestaurnatGateway(ctx, db)}
}

func (s *RestaurantHTTPService) Handler(w http.ResponseWriter, r *http.Request) {
	rts, err := s.gtw.ListRestaurants()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "<h1>Somethig Bad happened error: %s</h1>\n", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rts)
}
