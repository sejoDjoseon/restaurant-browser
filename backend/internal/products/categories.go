package products

import (
	"restaurant-browser/core/entities"
	"restaurant-browser/internal/money"
)

func parseDBProduct(prDB entities.DBProduct) entities.Product {
	return entities.Product{
		ID:          prDB.ID,
		Name:        prDB.Name,
		Description: prDB.Description,
		Image:       prDB.Image,
		Price: money.Value{
			Amount:   prDB.PriceValue,
			Currency: money.Currency(prDB.PriceCurrency),
		},
	}
}

func CategoriesFromDBProducts(DBproducts []entities.DBProduct) []entities.Category {
	categories := make(map[string]entities.Category)
	for _, prDB := range DBproducts {
		category, ok := categories[prDB.Category]
		if !ok {
			categories[prDB.Category] = entities.Category{
				Category: prDB.Category,
				Products: []entities.Product{
					parseDBProduct(prDB),
				},
			}
			continue
		}
		category.Products = append(category.Products, parseDBProduct(prDB))
		categories[prDB.Category] = category
	}

	result := []entities.Category{}
	for _, value := range categories {
		result = append(result, value)
	}

	return result
}
