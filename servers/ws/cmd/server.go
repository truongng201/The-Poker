package main

import (
	config "ws-service/config"
	controller "ws-service/internal/controller"
	routes "ws-service/internal/routes"
	utils "ws-service/pkg/utils"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
)

func main() {
	config.LoadConfig()

	e := echo.New()

	e.Validator = &utils.CustomValidator{Validator: validator.New()}
	e.Binder = &utils.CustomBinder{}

	controller := &controller.AppController{}

	e = routes.Routes(e, controller)

	switch config.Con.Environment {
	case "development":
		log.SetFormatter(&utils.CustomTextFormatter{})
		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format:           "timestamp=${time_rfc3339_nano} method=${method} remote_ip=${remote_ip} uri=${uri} status=${status} error=${error}\n",
			CustomTimeFormat: "2006-01-02 15:04:05",
		})) // timestamp=2023-08-23T15:49:32.565481601Z method=GET remote_ip=172.19.0.1 uri=/health status=200 error=
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: config.Con.AllowedOrigins,
		}))

		e.Logger.Info("Server is running on port 8080")
		e.Logger.Fatal(e.Start(":8080"))

	case "production":
		utils.CustomLogConfig()
		e.Use(middleware.Logger())
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: config.Con.AllowedOrigins,
		}))

		e.Logger.Info("Server is running on port 8080")
		e.Logger.Fatal(e.Start(":8080"))
	default:
		utils.CustomLogConfig()
		e.Use(middleware.Logger())
		e.Logger.Fatal("Environment not set")
		return
	}
}
