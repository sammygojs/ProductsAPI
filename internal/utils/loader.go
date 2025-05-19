package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"ProductsAPI/internal/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func LoadProductsFromDynamo() (*models.Products, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		return nil, fmt.Errorf("❌ Failed to load AWS config: %w", err)
	}

	client := dynamodb.NewFromConfig(cfg)

	out, err := client.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: aws.String("ProductsTable"),
	})
	if err != nil {
		return nil, fmt.Errorf("❌ Failed to scan DynamoDB: %w", err)
	}

	var productList []*models.Product
	if err := attributevalue.UnmarshalListOfMaps(out.Items, &productList); err != nil {
		return nil, fmt.Errorf("❌ Failed to unmarshal products: %w", err)
	}

	return &models.Products{
		Count:    len(productList),
		Total:    len(productList),
		Products: productList,
	}, nil
}
