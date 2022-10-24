package restaurant

import (
	"context"
	"database/sql"
	"restaurant-browser/core/entities"
	"restaurant-browser/internal/products"
)

type RestaurantGateway interface {
	ListRestaurants() ([]entities.Restaurant, error)
	RestaurantCatalog(rstID string) (entities.Catalog, error)
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

func (l *RestaurantLogic) RestaurantCatalog(rstID string) (entities.Catalog, error) {
	rstDBProducts, err := l.St.listRestaurantProductsInDB(rstID)
	if err != nil {
		return nil, err
	}

	// create categories
	categories := products.CategoriesFromDBProducts(rstDBProducts)

	return categories, nil
}
