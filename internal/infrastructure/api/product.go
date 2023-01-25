package infrastructure

import (
	"errors"
	"github.com/labstack/echo/v4"
	create "md-stock/internal/application/product/create"
	getAll "md-stock/internal/application/product/getAll"
	domain "md-stock/internal/domain/shared"
	infrastructure "md-stock/internal/infrastructure/product/model"
	shared "md-stock/internal/infrastructure/shared"
	"net/http"
	"strconv"
)

type ProductApi struct {
	createUseCase create.CreateProductUseCase
	getAllUseCase getAll.GetAllProductUseCase
}

func NewProductApi(createUseCase create.CreateProductUseCase, getAllUseCase getAll.GetAllProductUseCase) *ProductApi {
	return &ProductApi{
		createUseCase: createUseCase,
		getAllUseCase: getAllUseCase,
	}
}

func (api *ProductApi) Create(ctx echo.Context) error {
	var request infrastructure.CreateProductRequest

	if err := ctx.Bind(&request); err != nil {
		return err
	}

	command := create.NewCreateProductCommand(request.Name, request.Description, request.Price, request.Active)

	output, err := api.createUseCase.Execute(command)
	if err != nil {
		err := ctx.JSON(http.StatusBadRequest, shared.NewErrorResponse(err))
		return err
	}

	ctx.JSON(http.StatusCreated, infrastructure.NewCreateProductResponseFrom(output))

	return nil
}

func (api *ProductApi) GetAll(ctx echo.Context) error {
	page, err := strconv.Atoi(ctx.QueryParam("page"))
	if err != nil {
		return errors.New("query parameter 'page' should be an integer")
	}
	perPage, err := strconv.Atoi(ctx.QueryParam("perPage"))
	if err != nil {
		return errors.New("query parameter 'perPage' should be an integer")
	}

	searchQuery := domain.NewSearchQuery(page, perPage, nil, nil, nil)
	output, err := api.getAllUseCase.Execute(searchQuery)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, shared.NewErrorResponse(err))
		return err
	}

	ctx.JSON(http.StatusOK, infrastructure.NewGetAllProductResponseFrom(output))

	return nil
}
