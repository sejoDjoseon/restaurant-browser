package restaurant

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"restaurant-browser/core/entities"
	"restaurant-browser/internal/coordinates"
)

type RestaurantStorage interface {
	listRestaurantsInDB() ([]entities.Restaurant, error)
	listRestaurantProductsInDB(restID string) ([]entities.DBProduct, error)
}

type RestaurantService struct {
	db  *sql.DB
	ctx context.Context
}

func NewRestaurantStorage(ctx context.Context, db *sql.DB) RestaurantStorage {
	return &RestaurantService{ctx: ctx, db: db}
}

func (s *RestaurantService) listRestaurantsInDB() ([]entities.Restaurant, error) {
	querystring := `SELECT id, _name, _image, _location FROM t_restaurant`
	rows, err := s.db.QueryContext(s.ctx, querystring)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var restaurants []entities.Restaurant
	for rows.Next() {
		var rt entities.Restaurant
		var point []byte
		err = rows.Scan(&rt.ID, &rt.Name, &rt.Image, &point)
		if err != nil {
			return nil, err
		}
		rt.Location, err = coordinates.DecodePoint(point)
		if err != nil {
			fmt.Fprintf(os.Stdout, "error decoding coordinate:%s\n", err)
		}

		restaurants = append(restaurants, rt)
	}

	return restaurants, nil
}

func (s *RestaurantService) listRestaurantProductsInDB(restID string) ([]entities.DBProduct, error) {
	querystring := fmt.Sprintf(`SELECT id,
		restaurant_id,
		_category,
		_name,
		_description,
		_image,
		_price_value,
		_price_currency 
		FROM t_product
		WHERE restaurant_id = '%s'`, restID)

	rows, err := s.db.QueryContext(s.ctx, querystring)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []entities.DBProduct
	for rows.Next() {
		var pr entities.DBProduct
		err = rows.Scan(
			&pr.ID,
			&pr.RestaurantID,
			&pr.Category,
			&pr.Name,
			&pr.Description,
			&pr.Image,
			&pr.PriceValue,
			&pr.PriceCurrency,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, pr)
	}

	return products, nil
}
