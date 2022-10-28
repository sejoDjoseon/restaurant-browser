package restaurant

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"restaurant-browser/core/entities"
	"restaurant-browser/internal/coordinates"
	"restaurant-browser/internal/money"
	"restaurant-browser/internal/sqlmodels"
)

type RestaurantStorage interface {
	listRestaurants() ([]entities.Restaurant, error)
	listRestaurantProducts(restID string) ([]entities.Product, error)
}

type RestaurantService struct {
	db  *sql.DB
	ctx context.Context
}

func NewRestaurantStorage(ctx context.Context, db *sql.DB) RestaurantStorage {
	return &RestaurantService{ctx: ctx, db: db}
}

func (s *RestaurantService) listRestaurants() ([]entities.Restaurant, error) {
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

func parseDBProduct(prDB sqlmodels.DBProduct) entities.Product {
	return entities.Product{
		ID:           prDB.ID,
		RestaurantID: prDB.RestaurantID,
		Category:     prDB.Category,
		Name:         prDB.Name,
		Description:  prDB.Description,
		Image:        prDB.Image,
		Price: money.Value{
			Amount:   prDB.PriceValue,
			Currency: money.Currency(prDB.PriceCurrency),
		},
	}
}

func (s *RestaurantService) listRestaurantProducts(restID string) ([]entities.Product, error) {
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

	var products []entities.Product
	for rows.Next() {
		var pr sqlmodels.DBProduct
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
		products = append(products, parseDBProduct(pr))
	}

	return products, nil
}
