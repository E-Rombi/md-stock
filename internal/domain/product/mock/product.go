package domain

import (
	"errors"
	"github.com/stretchr/testify/mock"
	domain "md-stock/internal/domain/product"
	shared "md-stock/internal/domain/shared"
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
	product := domain.Product{
		ID:          "",
		Name:        "",
		Description: "",
		Price:       0,
		Active:      false,
		CreatedAt:   nil,
		UpdatedAt:   nil,
	}

	return shared.NewPagination[domain.Product](aQuery.Page, aQuery.PerPage, 0, []domain.Product{product}), nil
}
