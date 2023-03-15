package main

import (
	"github.com/labstack/echo/v4"
	infrastructure "md-stock/internal/infrastructure/configuration"
)

func main() {
	server := echo.New()

	infrastructure.NewApplication(server).Start()

	server.Logger.Fatal(server.Start(":8080"))
}
