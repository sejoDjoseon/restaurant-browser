package entities

import (
	"restaurant-browser/internal/coordinates"
	"restaurant-browser/internal/money"
)

type Restaurant struct {
	ID       string            `json:"id"`
	Name     string            `json:"name"`
	Image    string            `json:"image,omitempty"`
	Location coordinates.Point `json:"location"`
	Open     bool              `json:"open"`
}

type Product struct {
	ID           string
	RestaurantID string
	Category     string
	Name         string
	Description  string
	Image        string
	Price        money.Value
}

type CategoryProduct struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Image       string      `json:"image,omitempty"`
	Price       money.Value `json:"price"`
}

type Category struct {
	Category string            `json:"category"`
	Products []CategoryProduct `json:"products"`
}

type Catalog []Category
