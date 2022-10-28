package restaurant

import (
	"context"
	"database/sql"
	"restaurant-browser/core/entities"
	"restaurant-browser/internal/coordinates"
	"restaurant-browser/internal/products"
	"restaurant-browser/internal/restaurant_utils"
)

type RestaurantGateway interface {
	ListRestaurants(location *coordinates.Point) ([]entities.Restaurant, error)
	RestaurantCatalog(rstID string, filter *string) (entities.Catalog, error)
}

type RestaurantLogic struct {
	St        RestaurantStorage
	ScheduleS ScheduleService
}

func NewRestaurnatGateway(ctx context.Context, db *sql.DB) RestaurantGateway {
	return &RestaurantLogic{St: NewRestaurantStorage(ctx, db), ScheduleS: NewScheduleService()}
}

func (l *RestaurantLogic) ListRestaurants(location *coordinates.Point) ([]entities.Restaurant, error) {
	restaurants, err := l.St.listRestaurants()
	if err != nil {
		return restaurants, err
	}

	if location != nil {
		restaurants = restaurant_utils.SortByDistance(restaurants, *location)
	}

	for i := range restaurants {
		restaurants[i].Open = l.ScheduleS.isOpen(restaurants[i].ID)
	}

	return restaurants, nil
}

func (l *RestaurantLogic) RestaurantCatalog(rstID string, filter *string) (entities.Catalog, error) {
	rstProducts, err := l.St.listRestaurantProducts(rstID)
	if err != nil {
		return nil, err
	}

	if filter != nil {
		rstProducts = products.FilterProducts(rstProducts, *filter)
	}

	// create categories
	categories := products.CategoriesFromProducts(rstProducts)

	return categories, nil
}
