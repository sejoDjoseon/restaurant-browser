package restaurant

import (
	"context"
	"database/sql"
	"restaurant-browser/core/entities"
)

type RestaurantGateway interface {
	ListRestaurants() ([]entities.Restaurant, error)
}

type RestaurantLogic struct {
	St RestaurantStorage
}

func NewRestaurnatGateway(ctx context.Context, db *sql.DB) RestaurantGateway {
	return &RestaurantLogic{St: NewRestaurantStorage(ctx, db)}
}

func (l *RestaurantLogic) ListRestaurants() ([]entities.Restaurant, error) {
	return l.St.listRestaurantsInDB()
}
