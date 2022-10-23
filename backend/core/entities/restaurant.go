package entities

import "restaurant-browser/internal/coordinates"

type Restaurant struct {
	ID       string            `json:"id"`
	Name     string            `json:"name"`
	Image    string            `json:"image,omitempty"`
	Location coordinates.Point `json:"location"`
}
