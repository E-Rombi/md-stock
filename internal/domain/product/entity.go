package domain

import (
	"errors"
	"github.com/google/uuid"
	"md-stock/internal/domain/shared"
	"strings"
	"time"
)

type Product struct {
	ID          string
	Name        string
	Description string
	Price       domain.Money
	Active      bool
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}

func NewProduct(aName string, aDescription string, aPrice float64, isActive bool) (*Product, error) {
	id := uuid.NewString()
	price := domain.Money(aPrice * 100)
	now := time.Now()

	product := &Product{
		ID:          id,
		Name:        aName,
		Description: aDescription,
		Price:       price,
		Active:      isActive,
		CreatedAt:   &now,
	}

	if err := product.validate(); err != nil {
		return nil, err
	}

	return product, nil
}

func NewProductWith(anId string, aName string, aDescription string, aPrice int64, isActive bool, createdAt *time.Time, updatedAt *time.Time) (*Product, error) {
	price := domain.Money(aPrice * 100)

	product := &Product{
		ID:          anId,
		Name:        aName,
		Description: aDescription,
		Price:       price,
		Active:      isActive,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}

	if err := product.validate(); err != nil {
		return nil, err
	}

	return product, nil
}

func (p *Product) validate() error {
	var messages []string
	if strings.TrimSpace(p.Name) == "" {
		messages = append(messages, "'name' is required")
	}
	if strings.TrimSpace(p.Description) == "" {
		messages = append(messages, "'description' is required")
	}
	if p.Price <= 0 {
		messages = append(messages, "'price' is required and cannot be zero")
	}
	if len(messages) > 0 {
		return errors.New(strings.Join(messages, ","))
	}

	return nil
}

func (p *Product) inactivate() {
	now := time.Now()

	p.Active = false
	p.UpdatedAt = &now
}

func (p *Product) activate() {
	now := time.Now()

	p.Active = true
	p.UpdatedAt = &now
}
