package application

type CreateProductUseCase interface {
	execute(aCommand *CreateProductCommand) (*CreateProductOutput, error)
}
