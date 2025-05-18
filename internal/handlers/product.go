package handlers

import (
	"net/http"
	"strconv"
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

	idParam := c.Param("productID")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	for _, p := range cachedProducts.Products {
		if p.ID == id {
			c.JSON(http.StatusOK, p)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
}
