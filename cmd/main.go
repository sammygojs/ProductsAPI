package main

import (
	"github.com/gin-gonic/gin"
	"ProductsAPI/internal/handlers"
	"fmt"
)

func main() {
	router := gin.Default()

	api := router.Group("/api")
	{
		fmt.Println("✅ Registered /api/query-db route")
		api.GET("/query-db", handlers.QueryDB) // ✅ Correct group

		products := api.Group("/products")
		{
			products.GET("/", handlers.GetProducts)
			products.GET("/:productID", handlers.GetProduct)
		}
	}
	for _, route := range router.Routes() {
		fmt.Printf("📦 ROUTE REGISTERED: %s %s\n", route.Method, route.Path)
	}

	router.Run(":8080")
}
