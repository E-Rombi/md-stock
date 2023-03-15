package mock

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	domain "md-stock/internal/domain/product"
	shared "md-stock/internal/domain/shared"
	"time"
)

type MockProductGateway struct {
	mock.Mock
}

func (m *MockProductGateway) GetAll(aQuery *shared.SearchQuery) (*shared.Pagination[domain.Product], error) {
	id_1 := uuid.NewString()
	now := time.Now()

	someItems := []domain.Product{
		{
			id_1,
			"Product 1",
			"Product 1 description",
			3200,
			true,
			&now,
			nil,
		},
	}

	return shared.NewPagination[domain.Product](0, 10, 20, someItems), nil
}

func (g *MockProductGateway) Create(aProduct *domain.Product) (*domain.Product, error) {
	now := time.Now()
	id_1 := uuid.NewString()

	product, _ := domain.NewProductWith(id_1, "Product 1", "Product 1 description", 3200, true, &now, nil)

	return product, nil
}
