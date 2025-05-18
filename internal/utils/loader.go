package utils

import (
	"encoding/json"
	"log"
	"os"
	"ProductsAPI/internal/models"
)

func LoadProducts(fileName string) *models.Products {
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Failed to read Json file: %v", err)
	}

	var productList []*models.Product
	err = json.Unmarshal(data, &productList)
	if err != nil {
		log.Fatalf("Failed to unmarshall products: %v", err)
	}

	return &models.Products{
		Count:    len(productList),
		Total:    len(productList),
		Products: productList,
	}
}
