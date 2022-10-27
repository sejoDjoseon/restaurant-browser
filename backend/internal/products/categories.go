package products

import (
	"fmt"
	"os"
	"restaurant-browser/core/entities"
	"restaurant-browser/internal/money"
	"restaurant-browser/internal/string_utils"
	"strings"
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

func FilterProducts(products []entities.DBProduct, filter string) []entities.DBProduct {
	if filter == "" {
		return products
	}
	sanitizedFilter, err := sanitize(filter)
	if err != nil {
		return products
	}

	filteredProducts := []entities.DBProduct{}
	for _, pr := range products {
		if matchProduct(pr, sanitizedFilter) {
			filteredProducts = append(filteredProducts, pr)
		}
	}

	return filteredProducts
}

func sanitize(filter string) (string, error) {
	s, err := string_utils.RemoveAccents(filter)
	if err != nil {
		return "", err
	}

	return strings.ToLower(s), nil
}

func matchProduct(pr entities.DBProduct, filter string) bool {
	sName, err := sanitize(pr.Name)
	if err != nil {
		fmt.Fprintf(os.Stdout, "Wrong product if: %s\nError: %s\n", pr.ID, err.Error())
		return false
	}
	sDescription, err := sanitize(pr.Description)
	if err != nil {
		fmt.Fprintf(os.Stdout, "Wrong product if: %s\nError: %s\n", pr.ID, err.Error())
		return false
	}

	return strings.Contains(sName, filter) || strings.Contains(sDescription, filter)
}
