package application

import domain "md-stock/internal/domain/product"

type DefaultCreateProductUseCase struct {
	gateway domain.ProductGateway
}

func NewDefaultCreateProductUseCase(aGateway domain.ProductGateway) *DefaultCreateProductUseCase {
	return &DefaultCreateProductUseCase{
		gateway: aGateway,
	}
}

func (u *DefaultCreateProductUseCase) execute(aCommand *CreateProductCommand) (*CreateProductOutput, error) {
	name := aCommand.Name
	description := aCommand.Description
	price := aCommand.Price
	isActive := aCommand.Active

	product, err := domain.NewProduct(name, description, price, isActive)
	if err != nil {
		return nil, err
	}

	_, err = u.gateway.Create(product)
	if err != nil {
		return nil, err
	}

	return CreateProductOutputFrom(product), nil
}
