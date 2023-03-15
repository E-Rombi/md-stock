package application

import (
	"github.com/stretchr/testify/assert"
	domain "md-stock/internal/domain/product"
	mocking "md-stock/internal/domain/product/mock"
	shared "md-stock/internal/domain/shared"
	"testing"
)

func TestDefaultGetAllProductUseCase_Execute(t *testing.T) {
	gateway := mocking.NewProductGatewayMock()

	type fields struct {
		gateway domain.ProductGateway
	}
	type args struct {
		aQuery *shared.SearchQuery
	}
	type want struct {
		perPage int
		total   int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    want
		wantErr bool
	}{
		{
			name:   "should exists two products and return only one",
			fields: fields{gateway: gateway},
			args:   args{aQuery: shared.NewSearchQuery(0, 1, nil, nil, nil)},
			want: want{
				perPage: 1,
				total:   2,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &DefaultGetAllProductUseCase{
				gateway: tt.fields.gateway,
			}
			got, err := u.Execute(tt.args.aQuery)
			if (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.want.perPage, got.PerPage)
			assert.Equal(t, tt.want.total, got.Total)
		})
	}
}
