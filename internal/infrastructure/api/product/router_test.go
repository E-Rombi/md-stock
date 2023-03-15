package api

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	create "md-stock/internal/application/product/create"
	getall "md-stock/internal/application/product/getall"
	domain "md-stock/internal/domain/product"
	"md-stock/internal/infrastructure/persistence/product/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	mockProductGateway = mock.MockProductGateway{}
	product, _         = domain.NewProduct("Product 1", "Product 1 description", 32.0, true)
)

func TestProductApi_Create(t *testing.T) {
	json := getProductJson()

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(json))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	create := create.NewDefaultCreateProductUseCase(&mockProductGateway)
	get := &getall.DefaultGetAllProductUseCase{}

	api := NewProductApi(create, get)

	if assert.NoError(t, api.Create(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}

func getProductJson() string {
	bytes, _ := json.Marshal(product)

	return string(bytes)
}
