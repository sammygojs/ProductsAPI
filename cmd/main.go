package main

import (
	"github.com/gin-gonic/gin"
	"ProductsAPI/internal/handlers"
)

func main() {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("/query-db", handlers.QueryDB) // ✅ Correct group

		products := api.Group("/products")
		{
			products.GET("/", handlers.GetProducts)
			products.GET("/:productID", handlers.GetProduct)
		}
	}

	router.Run(":8080")
}
