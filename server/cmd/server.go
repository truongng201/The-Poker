package main

import (
	routes "server/pkg/routes"
	controller "server/pkg/controller"
	"github.com/labstack/echo/v4"
)

func main() {

	// Echo instance
	e := echo.New()

	controller := controller.AppController{}
	// Routes

	e = routes.Routes(e, controller)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}