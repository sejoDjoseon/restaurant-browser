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
}

type RestauranService struct {
	db  *sql.DB
	ctx context.Context
}

func NewRestaurantStorage(ctx context.Context, db *sql.DB) RestaurantStorage {
	return &RestauranService{ctx: ctx, db: db}
}

func (s *RestauranService) listRestaurantsInDB() ([]entities.Restaurant, error) {
	querystring := `SELECT id, _name, _image, _location FROM restaurant`
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
