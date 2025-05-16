package main

import (
	"net/http"

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
	router.Run(":3000")
}

func getProducts(c *gin.Context) {
	// Get products from JSON

	c.JSON(http.StatusOK, &Products{})
}

func getProduct(c *gin.Context) {
	// Get products from JSON based on a give id.

	c.JSON(http.StatusOK, &Product{})
}

type Products struct {
	Count    int         `json:"count"`
	Total    int         `json:"total"`
	Products []*Products `json:"products"`
}

type Product struct {
	ID               string    `json:"id"`
	SKU              string    `json:"sku"`
	ShortDescription *string   `json:"shortDescription"`
	LongDescription  *string   `json:"longDescription"`
	Brand            string    `json:"brand"`
	Price            *Money    `json:"price"`
	Features         []string  `json:"features,omitempty"`
	Images           []*Image  `json:"images"`
	Options          []*Option `json:"options"`
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
