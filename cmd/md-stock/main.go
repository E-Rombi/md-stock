package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	infrastructure "md-stock/internal/infrastructure/configuration"
)

func main() {
	server := echo.New()

	setUpMiddlewares(server)

	infrastructure.NewApplication(server).Start()

	server.Logger.Fatal(server.Start(":8080"))
}

func setUpMiddlewares(server *echo.Echo) {
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())
}
