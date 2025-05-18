package handlers

import (
	"fmt"
	"net/http"
	// "os"
	// "encoding/json"

	"github.com/gin-gonic/gin"
	"ProductsAPI/internal/models"
	"ProductsAPI/internal/utils"
)

var cachedProducts *models.Products

func SetCachedProducts(p *models.Products) {
	cachedProducts = p
}

func GetProducts(c *gin.Context) {
	if cachedProducts == nil {
		cachedProducts = utils.LoadProducts("products.json")
	}
	c.JSON(http.StatusOK, cachedProducts)
}

func GetProduct(c *gin.Context) {
	if cachedProducts == nil {
		cachedProducts = utils.LoadProducts("products.json")
	}
	id := c.Param("productID")
	for _, p := range cachedProducts.Products {
		if fmt.Sprintf("%v", p.ID) == id {
			c.JSON(http.StatusOK, p)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
}
