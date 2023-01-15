package application

import (
	"github.com/stretchr/testify/assert"
	domain "md-stock/internal/domain/product"
	mocking "md-stock/internal/domain/product/mock"
	"testing"
)

func TestDefaultCreateProductUseCase_execute(t *testing.T) {
	gateway := mocking.NewProductGatewayMock()

	type fields struct {
		gateway domain.ProductGateway
	}
	type args struct {
		aCommand *CreateProductCommand
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "Should create a product successfully",
			fields: fields{gateway: gateway},
			args: args{
				aCommand: NewCreateProductCommand(
					"Product 2",
					"Product 2 description",
					32.09,
					true,
				),
			},
			wantErr: false,
		},
		{
			name:   "Should throw an error when the property name is empty",
			fields: fields{gateway: gateway},
			args: args{
				aCommand: NewCreateProductCommand(
					"",
					"Product 1 description",
					32.09,
					true,
				),
			},
			wantErr: true,
		},
		{
			name:   "Should throw an error when gateway return an error too",
			fields: fields{gateway: gateway},
			args: args{
				aCommand: NewCreateProductCommand(
					"Product 1",
					"Product 1 description",
					32.09,
					true,
				),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &DefaultCreateProductUseCase{
				gateway: tt.fields.gateway,
			}
			got, err := uc.execute(tt.args.aCommand)
			if (err != nil) != tt.wantErr {
				t.Errorf("execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err == nil {
				assert.NotNil(t, got.ID)
				assert.NotEmpty(t, got.ID)
			}
		})
	}
}
