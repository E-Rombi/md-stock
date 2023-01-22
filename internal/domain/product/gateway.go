package domain

import domain "md-stock/internal/domain/shared"

type ProductGateway interface {
	Create(aProduct *Product) (*Product, error)
	GetAll(aQuery *domain.SearchQuery) (*domain.Pagination[Product], error)
}
