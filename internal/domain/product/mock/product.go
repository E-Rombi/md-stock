package domain

import (
	"errors"
	"github.com/stretchr/testify/mock"
	domain "md-stock/internal/domain/product"
	shared "md-stock/internal/domain/shared"
	"time"
)

type ProductGatewayMock struct {
	mock.Mock
}

func NewProductGatewayMock() *ProductGatewayMock {
	return &ProductGatewayMock{}
}

func (g *ProductGatewayMock) Create(aProduct *domain.Product) (*domain.Product, error) {
	if aProduct.Name == "Product 2" {
		return aProduct, nil
	}

	return nil, errors.New("id already exists")
}

func (g *ProductGatewayMock) GetAll(aQuery *shared.SearchQuery) (*shared.Pagination[domain.Product], error) {
	loc, _ := time.LoadLocation("America/Sao_Paulo")
	date := time.Date(2023, 1, 1, 1, 1, 1, 1, loc)

	products := []domain.Product{
		{
			ID:          "1",
			Name:        "Product 1",
			Description: "Product 1 description",
			Price:       3210,
			Active:      true,
			CreatedAt:   &date,
			UpdatedAt:   &date,
		},
		{
			ID:          "2",
			Name:        "Product 2",
			Description: "Product 2 description",
			Price:       3210,
			Active:      true,
			CreatedAt:   &date,
			UpdatedAt:   &date,
		},
	}

	if aQuery.PerPage == 1 {
		return shared.NewPagination[domain.Product](aQuery.Page, aQuery.PerPage, int64(len(products)), []domain.Product{products[0]}), nil
	} else {
		return shared.NewPagination[domain.Product](aQuery.Page, aQuery.PerPage, int64(len(products)), products), nil
	}

}
