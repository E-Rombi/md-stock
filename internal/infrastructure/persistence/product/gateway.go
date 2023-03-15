package persistence

import (
	"errors"
	"gorm.io/gorm"
	domain "md-stock/internal/domain/product"
	shared "md-stock/internal/domain/shared"
	"strings"
)

type ProductMySQLGateway struct {
	db *gorm.DB
}

func (g *ProductMySQLGateway) GetAll(aQuery *shared.SearchQuery) (*shared.Pagination[domain.Product], error) {
	offset := (aQuery.Page) * aQuery.PerPage

	where := buildWhere(aQuery.Terms)

	var totalItems int64
	results := g.db.
		Table("product").
		Where(where).
		Where(ProductGormEntity{Active: true}).
		Count(&totalItems)
	if results.Error != nil {
		return nil, errors.New("error during the query")
	}

	var entities []ProductGormEntity
	results = g.db.
		Offset(offset).
		Limit(aQuery.PerPage).
		Where(where).
		Where(ProductGormEntity{Active: true}).
		Find(&entities)

	if results.Error != nil {
		return nil, errors.New("error during the query")
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

func buildWhere(someTerms *string) map[string]string {
	if someTerms == nil || *someTerms == "" {
		return map[string]string{}
	}
	terms := strings.Split(*someTerms, ",")

	where := map[string]string{}
	for _, term := range terms {
		values := strings.Split(term, ":")
		where[values[0]] = values[1]
	}

	return where
}

func NewProductMySQLGateway(db *gorm.DB) *ProductMySQLGateway {
	return &ProductMySQLGateway{
		db: db,
	}
}

func (g *ProductMySQLGateway) Create(aProduct *domain.Product) (*domain.Product, error) {
	entity := NewProductGormEntityFrom(aProduct)

	err := g.db.Table("product").Create(entity).Error
	if err != nil {
		return nil, err
	}

	return aProduct, nil
}
