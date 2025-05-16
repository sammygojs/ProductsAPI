package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type MainSuite struct {
	suite.Suite
	router *gin.Engine
}

func (s *MainSuite) SetupTest() {
	s.router = gin.Default()

	s.router.GET(ROUTE_PRODUCTS, getProducts)
}

func (s *MainSuite) TestGetProducts() {
	req, err := http.NewRequest(http.MethodGet, ROUTE_PRODUCTS, nil)
	s.NoError(err)

	w := httptest.NewRecorder()

	s.router.ServeHTTP(w, req)

	s.Equal(http.StatusOK, w.Code)
	// @TODO: Add assertion of response
}

// @TODO: Add get product test

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestMainSuite(t *testing.T) {
	suite.Run(t, new(MainSuite))
}
