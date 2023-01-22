package application

import (
	domain "md-stock/internal/domain/product"
	"time"
)

type ProductListOutput struct {
	ID          string
	Name        string
	Description string
	Price       float64
	Active      bool
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}

func NewProductListOutputFrom(aProduct *domain.Product) *ProductListOutput {
	price := float64(aProduct.Price / 100)

	return &ProductListOutput{
		ID:          aProduct.ID,
		Name:        aProduct.Name,
		Description: aProduct.Description,
		Price:       price,
		Active:      aProduct.Active,
		CreatedAt:   aProduct.CreatedAt,
		UpdatedAt:   aProduct.UpdatedAt,
	}
}
