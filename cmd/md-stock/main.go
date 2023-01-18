package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	infrastructure "md-stock/internal/infrastructure/configuration"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	infrastructure.NewApplication(e).Start()

	e.Logger.Fatal(e.Start(":8080"))
}
