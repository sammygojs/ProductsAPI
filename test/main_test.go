package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"ProductsAPI/internal/handlers"
	"ProductsAPI/internal/models"
)

type MainSuite struct {
	suite.Suite
	router *gin.Engine
}

func (s *MainSuite) SetupTest() {
	gin.SetMode(gin.TestMode)

	// Set up test router
	s.router = gin.Default()
	s.router.GET("/api/products", handlers.GetProducts)

	// Inject mocked product data
	handlers.SetCachedProducts(&models.Products{
		Count: 1,
		Total: 1,
		Products: []*models.Product{
			{
				ID:   1,
				SKU:  "TEST123",
				Brand: "Test Brand",
			},
		},
	})
}

func (s *MainSuite) TestGetProducts() {
	req, err := http.NewRequest(http.MethodGet, "/api/products", nil)
	s.NoError(err)

	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, req)

	s.Equal(http.StatusOK, w.Code)

	var resp models.Products
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	s.NoError(err)

	s.Greater(resp.Count, 0, "Expected product count > 0")
	s.Greater(len(resp.Products), 0, "Expected at least one product in response")
	s.NotNil(resp.Products[0], "First product should not be nil")
	s.NotEmpty(resp.Products[0].SKU, "Product SKU should not be empty")
}

// To run the test suite
func TestMainSuite(t *testing.T) {
	suite.Run(t, new(MainSuite))
}
