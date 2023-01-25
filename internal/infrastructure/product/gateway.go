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
	offset := (aQuery.Page) * aQuery.PerPage

	var totalItems int64
	g.db.Table("product").Count(&totalItems)

	var entities []infrastructure.ProductGormEntity
	results := g.db.Offset(offset).Limit(aQuery.PerPage).Find(&entities)
	if results.Error != nil {
		return nil, results.Error
	}

	var products []domain.Product
	for _, entity := range entities {
		p, err := domain.NewProductWith(entity.ID, entity.Name, entity.Description, entity.Price, entity.Active, entity.CreatedAt, entity.UpdatedAt)
		if err != nil {
			return nil, err
		}

		products = append(products, *p)
	}

	return shared.NewPagination(aQuery.Page, aQuery.PerPage, totalItems, products), nil
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
