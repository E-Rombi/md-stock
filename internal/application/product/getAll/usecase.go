package application

import (
	domain "md-stock/internal/domain/product"
	shared "md-stock/internal/domain/shared"
)

type DefaultGetAllProductUseCase struct {
	gateway domain.ProductGateway
}

func (u *DefaultGetAllProductUseCase) Execute(aQuery *shared.SearchQuery) (*shared.Pagination[ProductListOutput], error) {
	page, err := u.gateway.GetAll(aQuery)
	if err != nil {
		return nil, err
	}

	output := make([]ProductListOutput, len(page.Items))
	for _, item := range page.Items {
		output = append(output, *NewProductListOutputFrom(&item))
	}

	return shared.NewPagination[ProductListOutput](page.CurrentPage, page.PerPage, 0, output), nil
}
