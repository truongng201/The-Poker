package main

import (
	controller "ws-service/pkg/controller"
	routes "ws-service/pkg/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	// Echo instance
	e := echo.New()
	controller := controller.AppController{}
	e = routes.Routes(e, controller)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))

}
