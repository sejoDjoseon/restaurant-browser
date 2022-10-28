package products

import (
	"restaurant-browser/core/entities"
	"restaurant-browser/internal/money"
	"testing"

	"github.com/stretchr/testify/assert"
)

type CategoriesFromProductsTest struct {
	ListProducts     []entities.Product
	CategoriesResult []entities.Category
}

var CategoriesFromProductsTests = []CategoriesFromProductsTest{
	{
		ListProducts:     []entities.Product{},
		CategoriesResult: []entities.Category{},
	},
	{ListProducts: []entities.Product{
		{
			ID:           "111111",
			RestaurantID: "111111",
			Category:     "1",
			Name:         "name",
			Description:  "description",
			Image:        "image",
			Price: money.Value{
				Amount:   1,
				Currency: money.EUR,
			},
		},
	},
		CategoriesResult: []entities.Category{
			{
				Category: "1",
				Products: []entities.CategoryProduct{
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
	{ListProducts: []entities.Product{
		{
			ID:           "111111",
			RestaurantID: "111111",
			Category:     "1",
			Name:         "name",
			Description:  "description",
			Image:        "image",
			Price: money.Value{
				Amount:   1,
				Currency: money.EUR,
			},
		},
		{
			ID:           "222222",
			RestaurantID: "111111",
			Category:     "1",
			Name:         "name",
			Description:  "description",
			Image:        "image",
			Price: money.Value{
				Amount:   1,
				Currency: money.EUR,
			},
		},
	},
		CategoriesResult: []entities.Category{
			{
				Category: "1",
				Products: []entities.CategoryProduct{
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
	{ListProducts: []entities.Product{
		{
			ID:           "111111",
			RestaurantID: "111111",
			Category:     "1",
			Name:         "name",
			Description:  "description",
			Image:        "image",
			Price: money.Value{
				Amount:   1,
				Currency: money.EUR,
			},
		},
		{
			ID:           "222222",
			RestaurantID: "111111",
			Category:     "1",
			Name:         "name",
			Description:  "description",
			Image:        "image",
			Price: money.Value{
				Amount:   1,
				Currency: money.EUR,
			},
		},
		{
			ID:           "333333",
			RestaurantID: "111111",
			Category:     "2",
			Name:         "name",
			Description:  "description",
			Image:        "image",
			Price: money.Value{
				Amount:   1,
				Currency: money.EUR,
			},
		},
	},
		CategoriesResult: []entities.Category{
			{
				Category: "1",
				Products: []entities.CategoryProduct{
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
				Products: []entities.CategoryProduct{
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

func TestCategoriesFromProducts(t *testing.T) {
	t.Parallel()
	for _, testCase := range CategoriesFromProductsTests {
		assert.Equal(t, testCase.CategoriesResult, CategoriesFromProducts(testCase.ListProducts))
	}
}

type FilterProductsTestCase struct {
	ListProducts []entities.Product
	Filter       string
	Result       []entities.Product
}

var FilterProductTestCases = []FilterProductsTestCase{
	{
		ListProducts: []entities.Product{},
		Filter:       "filter",
		Result:       []entities.Product{},
	},
	{
		ListProducts: []entities.Product{
			{
				ID:          "111111",
				Category:    "1",
				Name:        "name",
				Description: "description",
			},
		},
		Filter: "",
		Result: []entities.Product{
			{
				ID:          "111111",
				Category:    "1",
				Name:        "name",
				Description: "description",
			},
		},
	},
	{
		ListProducts: []entities.Product{
			{
				ID:          "111111",
				Category:    "1",
				Name:        "name",
				Description: "description",
			},
		},
		Filter: "filter",
		Result: []entities.Product{},
	},
	{
		ListProducts: []entities.Product{
			{
				ID:          "111111",
				Category:    "1",
				Name:        "name",
				Description: "description",
			},
		},
		Filter: "name",
		Result: []entities.Product{
			{
				ID:          "111111",
				Category:    "1",
				Name:        "name",
				Description: "description",
			},
		},
	},
	{
		ListProducts: []entities.Product{
			{
				ID:          "111111",
				Category:    "category",
				Name:        "name",
				Description: "description",
			},
		},
		Filter: "category",
		Result: []entities.Product{
			{
				ID:          "111111",
				Category:    "category",
				Name:        "name",
				Description: "description",
			},
		},
	},
	{
		ListProducts: []entities.Product{
			{
				ID:          "111111",
				Category:    "1",
				Name:        "name",
				Description: "descripción",
			},
		},
		Filter: "descripción",
		Result: []entities.Product{
			{
				ID:          "111111",
				Category:    "1",
				Name:        "name",
				Description: "descripción",
			},
		},
	},
	{
		ListProducts: []entities.Product{
			{
				ID:          "111111",
				Category:    "1",
				Name:        "polvorón",
				Description: "descripción",
			},
			{
				ID:          "2",
				Category:    "2",
				Name:        "no",
				Description: "descripción",
			},
			{
				ID:          "3",
				Category:    "2",
				Name:        "cesta navideña",
				Description: "con polvorones",
			},
			{
				ID:          "4",
				Category:    "polvorones",
				Name:        "POLVORON DE ALMENDRA",
				Description: "",
			},
			{
				ID:          "5",
				Category:    "polvorones",
				Name:        "Mantecado",
				Description: "",
			},
		},
		Filter: "polvorón",
		Result: []entities.Product{
			{
				ID:          "111111",
				Category:    "1",
				Name:        "polvorón",
				Description: "descripción",
			},
			{
				ID:          "3",
				Category:    "2",
				Name:        "cesta navideña",
				Description: "con polvorones",
			},
			{
				ID:          "4",
				Category:    "polvorones",
				Name:        "POLVORON DE ALMENDRA",
				Description: "",
			},
			{
				ID:          "5",
				Category:    "polvorones",
				Name:        "Mantecado",
				Description: "",
			},
		},
	},
}

func TestFilterProducts(t *testing.T) {
	t.Parallel()
	for _, tc := range FilterProductTestCases {
		assert.Equal(t, tc.Result, FilterProducts(tc.ListProducts, tc.Filter))
	}
}
