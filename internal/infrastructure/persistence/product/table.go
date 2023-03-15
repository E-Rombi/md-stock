package persistence

import (
	domain "md-stock/internal/domain/product"
	"time"
)

type ProductGormEntity struct {
	ID          string     `gorm:"column:id"`
	Name        string     `gorm:"column:name"`
	Description string     `gorm:"column:description"`
	Price       int64      `gorm:"column:price"`
	Active      bool       `gorm:"column:active"`
	CreatedAt   *time.Time `gorm:"column:created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at"`
}

func (ProductGormEntity) TableName() string {
	return "product"
}

func NewProductGormEntityFrom(aProduct *domain.Product) *ProductGormEntity {
	return &ProductGormEntity{
		ID:          aProduct.ID,
		Name:        aProduct.Name,
		Description: aProduct.Description,
		Price:       int64(aProduct.Price),
		Active:      aProduct.Active,
		CreatedAt:   aProduct.CreatedAt,
		UpdatedAt:   aProduct.UpdatedAt,
	}
}
