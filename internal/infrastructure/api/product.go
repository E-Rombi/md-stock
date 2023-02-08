package infrastructure

import (
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
	query, err := buildSearchQuery(ctx)
	if err != nil {
		return err
	}

	output, err := api.getAllUseCase.Execute(query)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, shared.NewErrorResponse(err))
		return err
	}

	ctx.JSON(http.StatusOK, infrastructure.NewGetAllProductResponseFrom(output))

	return nil
}

func buildSearchQuery(ctx echo.Context) (*domain.SearchQuery, error) {
	page, err := strconv.Atoi(ctx.QueryParam("page"))
	if err != nil {
		page = 0
	}
	perPage, err := strconv.Atoi(ctx.QueryParam("perPage"))
	if err != nil {
		perPage = 15
	}
	terms := ctx.QueryParam("terms")

	return domain.NewSearchQuery(page, perPage, &terms, nil, nil), nil
}
