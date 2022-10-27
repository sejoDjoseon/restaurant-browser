package restaurant_utils

import (
	"restaurant-browser/core/entities"
	"restaurant-browser/internal/coordinates"
	"sort"
)

func SortByDistance(restaurants []entities.Restaurant, point coordinates.Point) []entities.Restaurant {
	sort.SliceStable(restaurants, func(i, j int) bool {
		return coordinates.Distance(restaurants[i].Location, point) < coordinates.Distance(restaurants[j].Location, point)
	})
	return restaurants
}
