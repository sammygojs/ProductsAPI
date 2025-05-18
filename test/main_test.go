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

	s.router = gin.Default()
	s.router.GET("/api/products", handlers.GetProducts)
	s.router.GET("/api/products/:productID", handlers.GetProduct) // ✅ FIXED HERE

	// Inject rich mock product
	shortDesc := "Short Description"
	longDesc := "Long Description"
	handlers.SetCachedProducts(&models.Products{
		Count: 1,
		Total: 1,
		Products: []*models.Product{
			{
				ID:               1,
				SKU:              "ABC123",
				Brand:            "Nike",
				ShortDescription: &shortDesc,
				LongDescription:  &longDesc,
				Price: &models.Money{
					Amount:   "119.99",
					Currency: "GBP",
				},
				Images: []*models.Image{
					{Url: "https://example.com/img1.jpg"},
				},
				Variants: []*models.Variant{
					{
						ID: 11,
						EAN: "12345678901",
						SKU: "ABC123_7",
						Prices: struct {
							Price           float64     `json:"price"`
							MembershipPrice interface{} `json:"membershipPrice"`
							CurrencyCode    string      `json:"currencyCode"`
						}{
							Price:        119.99,
							MembershipPrice: 99.99,
							CurrencyCode: "GBP",
						},
						Inventory: struct {
							Count     interface{} `json:"count"`
							IsInStock bool        `json:"isInStock"`
						}{
							Count: nil, IsInStock: true,
						},
						Options: []struct {
							ID    int `json:"id"`
							Value []struct {
								Label        string `json:"label"`
								PreorderDate string `json:"preorderDate"`
							} `json:"value"`
							Group string `json:"group"`
						}{
							{
								ID: 1,
								Value: []struct {
									Label        string `json:"label"`
									PreorderDate string `json:"preorderDate"`
								}{
									{Label: "7", PreorderDate: ""},
								},
								Group: "Adult Footwear Size",
							},
						},
					},
				},
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

	s.Greater(resp.Count, 0)
	s.Greater(len(resp.Products), 0)
	s.NotNil(resp.Products[0])
	s.NotEmpty(resp.Products[0].SKU)
}

func (s *MainSuite) TestGetProductByID_StructureValidation() {
	req, _ := http.NewRequest(http.MethodGet, "/api/products/1", nil)
	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, req)

	s.Equal(http.StatusOK, w.Code)

	var product models.Product
	err := json.Unmarshal(w.Body.Bytes(), &product)
	s.NoError(err)

	s.Equal(1, product.ID)
	s.Equal("ABC123", product.SKU)
	s.NotNil(product.ShortDescription)
	s.NotEmpty(product.Images)

	s.Greater(len(product.Variants), 0)
	s.Equal("GBP", product.Variants[0].Prices.CurrencyCode)
	s.Greater(product.Variants[0].Prices.Price, 0.0)
}

// Optional: not found test
func (s *MainSuite) TestGetProductByID_NotFound() {
	req, _ := http.NewRequest(http.MethodGet, "/api/products/9999", nil)
	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, req)

	s.Equal(http.StatusNotFound, w.Code)

	var body map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &body)
	s.NoError(err)
	s.Equal("Product not found", body["error"])
}

func TestMainSuite(t *testing.T) {
	suite.Run(t, new(MainSuite))
}
