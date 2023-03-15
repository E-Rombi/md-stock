package api

import (
	"github.com/labstack/echo/v4"
	create "md-stock/internal/application/product/create"
	getAll "md-stock/internal/application/product/getall"
	domain "md-stock/internal/domain/shared"
	shared "md-stock/internal/infrastructure/api"
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

func (api *ProductApi) Register(server *echo.Echo) {
	server.POST("/products", api.Create)
	server.GET("/products", api.GetAll)
}

func (api *ProductApi) Create(ctx echo.Context) error {
	var request CreateProductRequest
	if err := ctx.Bind(&request); err != nil {
		return err
	}

	command := create.NewCreateProductCommand(request.Name, request.Description, request.Price, request.Active)

	output, err := api.createUseCase.Execute(command)
	if err != nil {
		err := ctx.JSON(http.StatusBadRequest, shared.NewErrorResponse(err))
		return err
	}

	ctx.JSON(http.StatusCreated, NewCreateProductResponseFrom(output))

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

	ctx.JSON(http.StatusOK, NewGetAllProductResponseFrom(output))

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
