package domain

import (
	"errors"
	"github.com/stretchr/testify/mock"
	domain "md-stock/internal/domain/product"
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
