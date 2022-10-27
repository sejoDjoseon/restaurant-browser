package restaurant_utils

import (
	"restaurant-browser/core/entities"
	"restaurant-browser/internal/coordinates"
	"testing"

	"github.com/stretchr/testify/assert"
)

type SortByDistanceTestCase struct {
	Restaurants []entities.Restaurant
	Point       coordinates.Point
	Result      []entities.Restaurant
}

var testCases []SortByDistanceTestCase = []SortByDistanceTestCase{
	{
		Restaurants: []entities.Restaurant{},
		Point: coordinates.Point{
			X: 0,
			Y: 0,
		},
		Result: []entities.Restaurant{},
	},
	{
		Restaurants: []entities.Restaurant{
			{
				Location: coordinates.Point{X: 0, Y: 0},
			},
		},
		Point: coordinates.Point{
			X: 0,
			Y: 0,
		},
		Result: []entities.Restaurant{
			{
				Location: coordinates.Point{X: 0, Y: 0},
			},
		},
	},
	{
		Restaurants: []entities.Restaurant{
			{
				ID:       "1",
				Location: coordinates.Point{X: 0, Y: 0},
			},
			{
				ID:       "2",
				Location: coordinates.Point{X: 0, Y: 0},
			},
		},
		Point: coordinates.Point{
			X: 0,
			Y: 0,
		},
		Result: []entities.Restaurant{
			{
				ID:       "1",
				Location: coordinates.Point{X: 0, Y: 0},
			},
			{
				ID:       "2",
				Location: coordinates.Point{X: 0, Y: 0},
			},
		},
	},
	{
		Restaurants: []entities.Restaurant{
			{
				ID:       "1",
				Location: coordinates.Point{X: 1, Y: 1},
			},
			{
				ID:       "2",
				Location: coordinates.Point{X: 1, Y: 0},
			},
			{
				ID:       "3",
				Location: coordinates.Point{X: 3, Y: 3},
			},
			{
				ID:       "4",
				Location: coordinates.Point{X: -2, Y: -2},
			},
		},
		Point: coordinates.Point{
			X: 0,
			Y: 0,
		},
		Result: []entities.Restaurant{
			{
				ID:       "2",
				Location: coordinates.Point{X: 1, Y: 0},
			},
			{
				ID:       "1",
				Location: coordinates.Point{X: 1, Y: 1},
			},
			{
				ID:       "4",
				Location: coordinates.Point{X: -2, Y: -2},
			},
			{
				ID:       "3",
				Location: coordinates.Point{X: 3, Y: 3},
			},
		},
	},
}

func TestCategoriesFromDBProducts(t *testing.T) {
	t.Parallel()
	for _, testCase := range testCases {
		assert.ElementsMatch(t, testCase.Result, SortByDistance(testCase.Restaurants, testCase.Point))
	}
}
