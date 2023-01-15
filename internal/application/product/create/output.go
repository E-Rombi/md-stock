package application

import domain "md-stock/internal/domain/product"

type CreateProductOutput struct {
	ID string
}

func CreateProductOutputFrom(aProduct *domain.Product) *CreateProductOutput {
	return &CreateProductOutput{
		ID: aProduct.ID,
	}
}
