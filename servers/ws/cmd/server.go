package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	
	// Echo instance
	e := echo.New()

	// Static Routes
	// ApiRoutes
	api := e.Group("/api/v1")

	api.GET("/health", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	// Start server
	e.Logger.Fatal(e.Start(":8080"))	

}