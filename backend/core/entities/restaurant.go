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

type DBProduct struct {
	ID            string
	RestaurantID  string
	Category      string
	Name          string
	Description   string
	Image         string
	PriceValue    int
	PriceCurrency string
}

type Product struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Image       string      `json:"image,omitempty"`
	Price       money.Value `json:"price"`
}

type Category struct {
	Category string    `json:"category"`
	Products []Product `json:"products"`
}

type Catalog []Category
