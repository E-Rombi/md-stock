package application

import (
	domain "md-stock/internal/domain/product"
	shared "md-stock/internal/domain/shared"
)

type DefaultGetAllProductUseCase struct {
	gateway domain.ProductGateway
}

func NewDefaultGetAllProductUseCase(gateway domain.ProductGateway) *DefaultGetAllProductUseCase {
	return &DefaultGetAllProductUseCase{
		gateway: gateway,
	}
}

func (u *DefaultGetAllProductUseCase) Execute(aQuery *shared.SearchQuery) (*shared.Pagination[ProductListOutput], error) {
	page, err := u.gateway.GetAll(aQuery)
	if err != nil {
		return nil, err
	}

	var output []ProductListOutput
	for _, item := range page.Items {
		output = append(output, *NewProductListOutputFrom(&item))
	}

	return shared.NewPagination[ProductListOutput](page.CurrentPage, page.PerPage, page.Total, output), nil
}
