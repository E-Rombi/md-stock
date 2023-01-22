package application

import domain "md-stock/internal/domain/shared"

type GetAllProductUseCase interface {
	Execute(aQuery *domain.SearchQuery) (*domain.Pagination[ProductListOutput], error)
}
