package application

type CreateProductUseCase interface {
	Execute(aCommand *CreateProductCommand) (*CreateProductOutput, error)
}
