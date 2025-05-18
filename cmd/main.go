package main

import (
	"github.com/gin-gonic/gin"
	"ProductsAPI/internal/handlers"
)

func main() {
	router := gin.Default()
	api := router.Group("/api")
	{
		p := api.Group("/products")
		{
			p.GET("/", handlers.GetProducts)
			p.GET("/:productID", handlers.GetProduct)
		}
	}
	router.Run(":8080")
}
