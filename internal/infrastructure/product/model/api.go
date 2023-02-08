package infrastructure

import (
	create "md-stock/internal/application/product/create"
	getAll "md-stock/internal/application/product/getAll"
	domain "md-stock/internal/domain/shared"
	"time"
)

type CreateProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Active      bool    `json:"active"`
}

type CreateProductResponse struct {
	ID string `json:"id"`
}

func NewCreateProductResponseFrom(anOutput *create.CreateProductOutput) *CreateProductResponse {
	return &CreateProductResponse{
		ID: anOutput.ID,
	}
}

type GetAllProductResponse struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Price       float64    `json:"price"`
	Active      bool       `json:"isActive"`
	CreatedAt   *time.Time `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt"`
}

func NewGetAllProductResponseFrom(anOutput *domain.Pagination[getAll.ProductListOutput]) *domain.Pagination[GetAllProductResponse] {
	var items []GetAllProductResponse
	for _, product := range anOutput.Items {
		items = append(items, GetAllProductResponse{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price / 100,
			Active:      product.Active,
			CreatedAt:   product.CreatedAt,
			UpdatedAt:   product.UpdatedAt,
		})

	}

	return domain.NewPagination(anOutput.CurrentPage, anOutput.PerPage, anOutput.Total, items)
}
