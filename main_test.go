package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type MainSuite struct {
	suite.Suite
	router *gin.Engine
}

// const ROUTE_PRODUCTS = "/api/products"
// const ROUTE_PRODUCT = "/api/products/1"

func (s *MainSuite) SetupTest() {
	s.router = gin.Default()

	cachedProducts = &Products{
		Count: 1,
		Total: 1,
		Products: []*Product{
			{ID: 1, SKU: "TEST123", Brand: "Test Brand"},
		},
	}
	
	s.router.GET("/api/products", getProducts)

	// s.router.GET(ROUTE_PRODUCT, getProduct)
}

func (s *MainSuite) TestGetProducts() {
	req, err := http.NewRequest(http.MethodGet, "/api/products", nil)
	s.NoError(err)

	w := httptest.NewRecorder()
	// s.router.GET("/api/products", getProducts)
	s.router.ServeHTTP(w, req)

	s.Equal(http.StatusOK, w.Code)
	// @TODO: Add assertion of response
	var resp Products
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	s.NoError(err)

	// Fail the test if no products were returned
	s.Greater(resp.Count, 0, "Expected product count > 0")
	s.Greater(len(resp.Products), 0, "Expected at least one product in response")
	s.NotNil(resp.Products[0], "First product should not be nil")
	s.NotEmpty(resp.Products[0].SKU, "Product SKU should not be empty")
}

// @TODO: Add get product test
// func (s *MainSuite) TestGetProduct() {
// 	// Add a dummy ID that exists in your test dataset
// 	productID := "1"
// 	req, err := http.NewRequest(http.MethodGet, "/api/products/"+productID, nil)
// 	s.NoError(err)

// 	w := httptest.NewRecorder()
// 	// s.router.GET(ROUTE_PRODUCT, getProduct)
// 	s.router.ServeHTTP(w, req)

// 	s.Equal(http.StatusOK, w.Code)
// }
// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestMainSuite(t *testing.T) {
	suite.Run(t, new(MainSuite))
}
