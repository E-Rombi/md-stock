package domain

type ProductGateway interface {
	Create(aProduct *Product) (*Product, error)
}
