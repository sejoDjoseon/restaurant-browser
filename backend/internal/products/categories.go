package products

import (
	"fmt"
	"os"
	"restaurant-browser/core/entities"
	"restaurant-browser/internal/string_utils"
	"strings"
)

func productToCategoryProduct(pr entities.Product) entities.CategoryProduct {
	return entities.CategoryProduct{
		ID:          pr.ID,
		Name:        pr.Name,
		Description: pr.Description,
		Image:       pr.Image,
		Price:       pr.Price,
	}
}

func CategoriesFromProducts(products []entities.Product) []entities.Category {
	categories := make(map[string]entities.Category)
	for _, pr := range products {
		category, ok := categories[pr.Category]
		if !ok {
			categories[pr.Category] = entities.Category{
				Category: pr.Category,
				Products: []entities.CategoryProduct{
					productToCategoryProduct(pr),
				},
			}
			continue
		}
		category.Products = append(category.Products, productToCategoryProduct(pr))
		categories[pr.Category] = category
	}

	result := []entities.Category{}
	for _, value := range categories {
		result = append(result, value)
	}

	return result
}

func FilterProducts(products []entities.Product, filter string) []entities.Product {
	if filter == "" {
		return products
	}
	sanitizedFilter, err := sanitize(filter)
	if err != nil {
		return products
	}

	filteredProducts := []entities.Product{}
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

func matchProduct(pr entities.Product, filter string) bool {
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
	sCategory, err := sanitize(pr.Category)
	if err != nil {
		fmt.Fprintf(os.Stdout, "Wrong product if: %s\nError: %s\n", pr.ID, err.Error())
		return false
	}

	return strings.Contains(sName, filter) || strings.Contains(sDescription, filter) || strings.Contains(sCategory, filter)
}
