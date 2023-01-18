package infrastructure

import application "md-stock/internal/application/product/create"

type CreateProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Active      bool    `json:"active"`
}

type CreateProductResponse struct {
	ID string `json:"id"`
}

func NewCreateProductResponseFrom(anOutput *application.CreateProductOutput) *CreateProductResponse {
	return &CreateProductResponse{
		ID: anOutput.ID,
	}
}
