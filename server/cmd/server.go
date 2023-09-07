package main

import (
	"github.com/labstack/echo/v4"
)

func main() {

	// Echo instance
	e := echo.New()

	// Routes
	e.GET("/health", func(c echo.Context) error {
		return c.String(200, "Hello, World 1!")
	})

	e.Static("/", "../../client/build")

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}