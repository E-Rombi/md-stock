package infrastructure

type CreateProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Active      bool    `json:"active"`
}

type CreateProductResponse struct {
	ID string `json:"id"`
}
