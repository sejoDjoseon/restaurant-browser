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
	restaurants, err := l.St.listRestaurantsInDB()
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
	rstDBProducts, err := l.St.listRestaurantProductsInDB(rstID)
	if err != nil {
		return nil, err
	}

	if filter != nil {
		rstDBProducts = products.FilterProducts(rstDBProducts, *filter)
	}

	// create categories
	categories := products.CategoriesFromDBProducts(rstDBProducts)

	return categories, nil
}
