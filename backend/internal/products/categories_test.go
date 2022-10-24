package products

import (
	"restaurant-browser/core/entities"
	"restaurant-browser/internal/money"
	"testing"

	"github.com/stretchr/testify/assert"
)

type CategoriesFromDBProductsTest struct {
	ListProducts     []entities.DBProduct
	CategoriesResult []entities.Category
}

var CategoriesFromDBProductsTests = []CategoriesFromDBProductsTest{
	{
		ListProducts:     []entities.DBProduct{},
		CategoriesResult: []entities.Category{},
	},
	{ListProducts: []entities.DBProduct{
		{
			ID:            "111111",
			RestaurantID:  "111111",
			Category:      "1",
			Name:          "name",
			Description:   "description",
			Image:         "image",
			PriceValue:    1,
			PriceCurrency: "EUR",
		},
	},
		CategoriesResult: []entities.Category{
			{
				Category: "1",
				Products: []entities.Product{
					{
						ID:          "111111",
						Name:        "name",
						Description: "description",
						Image:       "image",
						Price: money.Value{
							Amount:   1,
							Currency: money.EUR,
						},
					},
				},
			},
		},
	},
	{ListProducts: []entities.DBProduct{
		{
			ID:            "111111",
			RestaurantID:  "111111",
			Category:      "1",
			Name:          "name",
			Description:   "description",
			Image:         "image",
			PriceValue:    1,
			PriceCurrency: "EUR",
		},
		{
			ID:            "222222",
			RestaurantID:  "111111",
			Category:      "1",
			Name:          "name",
			Description:   "description",
			Image:         "image",
			PriceValue:    1,
			PriceCurrency: "EUR",
		},
	},
		CategoriesResult: []entities.Category{
			{
				Category: "1",
				Products: []entities.Product{
					{
						ID:          "111111",
						Name:        "name",
						Description: "description",
						Image:       "image",
						Price: money.Value{
							Amount:   1,
							Currency: money.EUR,
						},
					},
					{
						ID:          "222222",
						Name:        "name",
						Description: "description",
						Image:       "image",
						Price: money.Value{
							Amount:   1,
							Currency: money.EUR,
						},
					},
				},
			},
		},
	},
	{ListProducts: []entities.DBProduct{
		{
			ID:            "111111",
			RestaurantID:  "111111",
			Category:      "1",
			Name:          "name",
			Description:   "description",
			Image:         "image",
			PriceValue:    1,
			PriceCurrency: "EUR",
		},
		{
			ID:            "222222",
			RestaurantID:  "111111",
			Category:      "1",
			Name:          "name",
			Description:   "description",
			Image:         "image",
			PriceValue:    1,
			PriceCurrency: "EUR",
		},
		{
			ID:            "333333",
			RestaurantID:  "111111",
			Category:      "2",
			Name:          "name",
			Description:   "description",
			Image:         "image",
			PriceValue:    1,
			PriceCurrency: "EUR",
		},
	},
		CategoriesResult: []entities.Category{
			{
				Category: "1",
				Products: []entities.Product{
					{
						ID:          "111111",
						Name:        "name",
						Description: "description",
						Image:       "image",
						Price: money.Value{
							Amount:   1,
							Currency: money.EUR,
						},
					},
					{
						ID:          "222222",
						Name:        "name",
						Description: "description",
						Image:       "image",
						Price: money.Value{
							Amount:   1,
							Currency: money.EUR,
						},
					},
				},
			},
			{
				Category: "2",
				Products: []entities.Product{
					{
						ID:          "333333",
						Name:        "name",
						Description: "description",
						Image:       "image",
						Price: money.Value{
							Amount:   1,
							Currency: money.EUR,
						},
					},
				},
			},
		},
	},
}

func TestCategoriesFromDBProducts(t *testing.T) {
	for _, testCase := range CategoriesFromDBProductsTests {
		assert.Equal(t, testCase.CategoriesResult, CategoriesFromDBProducts(testCase.ListProducts))
	}
}
