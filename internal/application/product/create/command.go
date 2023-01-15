package application

type CreateProductCommand struct {
	Name        string
	Description string
	Price       float64
	Active      bool
}

func NewCreateProductCommand(aName string, aDescription string, aPrice float64, isActive bool) *CreateProductCommand {
	return &CreateProductCommand{
		Name:        aName,
		Description: aDescription,
		Price:       aPrice,
		Active:      isActive,
	}
}
