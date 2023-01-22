package infrastructure

import (
	"gorm.io/gorm"
	domain "md-stock/internal/domain/product"
	shared "md-stock/internal/domain/shared"
	infrastructure "md-stock/internal/infrastructure/product/model"
)

type ProductMySQLGateway struct {
	db *gorm.DB
}

func (g *ProductMySQLGateway) GetAll(aQuery *shared.SearchQuery) (*shared.Pagination[domain.Product], error) {
	//TODO implement me
	panic("implement me")
}

func NewProductMySQLGateway(db *gorm.DB) *ProductMySQLGateway {
	return &ProductMySQLGateway{
		db: db,
	}
}

func (g *ProductMySQLGateway) Create(aProduct *domain.Product) (*domain.Product, error) {
	entity := infrastructure.NewProductGormEntityFrom(aProduct)

	err := g.db.Table("product").Create(entity).Error
	if err != nil {
		return nil, err
	}

	return aProduct, nil
}
