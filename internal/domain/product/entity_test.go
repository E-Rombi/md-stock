package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewProduct(t *testing.T) {
	type args struct {
		aName        string
		aDescription string
		aPrice       float64
		isActive     bool
	}
	tests := []struct {
		name    string
		args    args
		want    *Product
		wantErr bool
	}{
		{
			name: "Should create a product successfully",
			args: args{
				aName:        "Product 1",
				aDescription: "Product 1 description",
				aPrice:       32.10,
				isActive:     true,
			},
			want: &Product{
				Name:        "Product 1",
				Description: "Product 1 description",
				Price:       3210,
				Active:      true,
			},
			wantErr: false,
		},
		{
			name: "Should throw an error when Name is empty",
			args: args{
				aDescription: "Product 1 description",
				aPrice:       32.10,
				isActive:     true,
			},
			wantErr: true,
		},
		{
			name: "Should throw an error when Description is empty",
			args: args{
				aName:    "Product 1",
				aPrice:   32.10,
				isActive: true,
			},
			wantErr: true,
		},
		{
			name: "Should throw an error when Price is zero",
			args: args{
				aName:        "Product 1",
				aDescription: "Product 1 description",
				isActive:     true,
			},
			wantErr: true,
		},
		{
			name: "Should throw an error when Price is lower than zero",
			args: args{
				aName:        "Product 1",
				aDescription: "Product 1 description",
				aPrice:       -2,
				isActive:     true,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewProduct(tt.args.aName, tt.args.aDescription, tt.args.aPrice, tt.args.isActive)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err == nil {
				assert.Equal(t, tt.want.Name, got.Name)
				assert.Equal(t, tt.want.Description, got.Description)
				assert.Equal(t, tt.want.Price, got.Price)
				assert.Equal(t, tt.want.Active, got.Active)
				assert.NotNil(t, got.CreatedAt)
				assert.Nil(t, got.UpdatedAt)
			}
		})
	}
}

func TestProduct_inactivate(t *testing.T) {
	product, _ := NewProduct("Product 1", "Product 1 description", 32.11, true)

	assert.True(t, product.Active)
	assert.Nil(t, product.UpdatedAt)

	product.inactivate()

	assert.False(t, product.Active)
	assert.NotNil(t, product.UpdatedAt)
}

func TestProduct_activate(t *testing.T) {
	product, _ := NewProduct("Product 1", "Product 1 description", 32.11, true)

	assert.True(t, product.Active)
	assert.Nil(t, product.UpdatedAt)

	product.inactivate()

	momentAfterInactivate := product.UpdatedAt
	assert.False(t, product.Active)
	assert.NotNil(t, momentAfterInactivate)

	product.activate()

	assert.True(t, product.Active)
	assert.True(t, product.UpdatedAt.After(*momentAfterInactivate))
}
