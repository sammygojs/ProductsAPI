package main

import (
	"net/http"
	"encoding/json"
	"log"
	"os"
	"fmt"

	"github.com/gin-gonic/gin"
)

const (
	ROUTE_PRODUCTS = "/products"
)

func main() {
	// Set the router.
	router := gin.Default()

	// Setup route group for the API.
	api := router.Group("/api")
	{
		b := api.Group(ROUTE_PRODUCTS)
		{
			b.GET("/", getProducts)
			b.GET("/:productID", getProduct)
		}
	}

	// Start and run the server
	router.Run(":8080")
}

var cachedProducts *Products

type ImageList []*Image

func (il *ImageList) UnmarshalJSON(data []byte) error {
	var urls []string
	if err := json.Unmarshal(data, &urls); err != nil {
		return err
	}

	var result []*Image
	for _, u := range urls {
		result = append(result, &Image{
			Url:     u,
			AltText: nil, // Optional: derive from filename
		})
	}
	*il = result
	return nil
}

func loadProducts(fileName string) *Products{
	
	data, err := os.ReadFile(fileName)
	if err!=nil{
		log.Fatalf("Failed to read Json file: %v",err)
	}
	var productList []*Product
	err = json.Unmarshal(data,&productList)
	if err != nil{
		log.Fatalf("Failed to unmarshall products: %v", err)
	}
	return &Products{
		Count: len(productList),
		Total: len(productList),
		Products: productList,
	}
}

func getProducts(c *gin.Context) {
	// lazy load
	if cachedProducts == nil{
		cachedProducts = loadProducts("products.json")
	}
	
	c.JSON(http.StatusOK, cachedProducts)
	// c.JSON(http.StatusOK, &Products{})
}

func getProduct(c *gin.Context) {
	productID := c.Param("productID")

	// Optional: lazy-load if not already loaded
	if cachedProducts == nil {
		cachedProducts = loadProducts("products.json")
	}

	for _, p := range cachedProducts.Products {
		if fmt.Sprintf("%v", p.ID) == productID {
			c.JSON(http.StatusOK, p)
			return
		}
	}

	// Not found
	c.JSON(http.StatusNotFound, gin.H{
		"error": "Product not found",
	})
}

type Products struct {
	Count    int         `json:"count"`
	Total    int         `json:"total"`
	// Products []*Products `json:"products"`
	Products []*Product `json:"products"`
}

type Product struct {
	ID               int    `json:"id"`
	// ID               string    `json:"id"`
	SKU              string    `json:"sku"`
	ShortDescription *string   `json:"shortDescription"`
	LongDescription  *string   `json:"longDescription"`
	Brand            string    `json:"brand"`
	Price            *Money    `json:"price"`
	Features         []string  `json:"features,omitempty"`
	Images 			 ImageList `json:"images"`
	// Images           []*Image  `json:"images"`
	Options          []*Option `json:"options"`
	Variants         []*Variant  `json:"variants"`
}

type Variant struct {
	ID    int    `json:"id"`
	EAN   string `json:"ean"`
	SKU   string `json:"sku"`
	Prices struct {
		Price           float64     `json:"price"`
		MembershipPrice interface{} `json:"membershipPrice"` // float or empty string/null
		CurrencyCode    string      `json:"currencyCode"`
	} `json:"prices"`
	Inventory struct {
		Count     interface{} `json:"count"` // Sometimes null
		IsInStock bool        `json:"isInStock"`
	} `json:"inventory"`
	Options []struct {
		ID    int `json:"id"`
		Value []struct {
			Label        string `json:"label"`
			PreorderDate string `json:"preorderDate"`
		} `json:"value"`
		Group string `json:"group"`
	} `json:"options"`
}


type Money struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type Image struct {
	Url     string  `json:"url"`
	AltText *string `json:"altText"`
}

type Option struct {
	ID      string `json:"id"`
	SKU     string `json:"sku"`
	Price   *Money `json:"price"`
	InStock bool   `json:"inStock"`
	Size    string `json:"size"`
	Colour  string `json:"colour"`
}
