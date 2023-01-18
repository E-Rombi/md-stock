package infrastructure

import (
	"github.com/labstack/echo/v4"
	application "md-stock/internal/application/product/create"
	infrastructure "md-stock/internal/infrastructure/product/model"
	"net/http"
)

type ProductApi struct {
	createUseCase application.CreateProductUseCase
}

func NewProductApi(createUseCase application.CreateProductUseCase) *ProductApi {
	return &ProductApi{
		createUseCase: createUseCase,
	}
}

func (api *ProductApi) Create(ctx echo.Context) error {
	var request infrastructure.CreateProductRequest
	if err := ctx.Bind(&request); err != nil {
		return err
	}
	command := application.NewCreateProductCommand(request.Name, request.Description, request.Price, request.Active)

	output, err := api.createUseCase.Execute(command)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}

	ctx.JSON(http.StatusCreated, output.ID)

	return nil
}
