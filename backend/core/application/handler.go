package app_restaurant

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	restaurant "restaurant-browser/core/infrastructure"

	"github.com/gorilla/mux"
)

type RestaurantHTTPService struct {
	gtw restaurant.RestaurantGateway
}

func NewRestaurantHTTPService(ctx context.Context, db *sql.DB) *RestaurantHTTPService {
	return &RestaurantHTTPService{restaurant.NewRestaurnatGateway(ctx, db)}
}

func (s *RestaurantHTTPService) RestaurantsListHandler(w http.ResponseWriter, r *http.Request) {
	rts, err := s.gtw.ListRestaurants()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, rts)
}

func (s *RestaurantHTTPService) RestaurantCatalogHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	catalog, err := s.gtw.RestaurantCatalog(id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, catalog)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
