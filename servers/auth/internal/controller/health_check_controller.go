package controller

import (
	"fmt"
	"runtime/debug"

	database "auth-service/pkg/database"
	utils "auth-service/pkg/utils"

	"github.com/labstack/echo/v4"
)

type HealthCheckController struct{}

type HealthCheck struct {
	Success bool   `json:"success" xml:"success"`
	Message string `json:"message" xml:"message"`
	Version string `json:"version" xml:"version"`
}

var Version string

var module_name = func() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		return info.Main.Path
	}
	return ""
}

func (controller *HealthCheckController) checkCacheConnection(c echo.Context) (bool, error) {
	_, err := utils.RedisClient.Ping(c.Request().Context()).Result()
	if err != nil {
		return false, c.JSON(500, &HealthCheck{
			Success: false,
			Message: "Cannot connect to cache",
			Version: Version,
		})
	}
	return true, nil
}

func (controller *HealthCheckController) checkDatabaseConnection(
	c echo.Context,
	store database.Store,
) (bool, error) {
	_, err := store.HealthCheck(c.Request().Context())
	if err != nil {
		return false, c.JSON(500, &HealthCheck{
			Success: false,
			Message: "Cannot connect to database",
			Version: Version,
		})
	}
	return true, nil
}

func (controller *HealthCheckController) Execute(c echo.Context, store database.Store) error {
	if Version == "" {
		Version = "development"
	} else {
		Version = Version[:7]
	}

	ok, err := controller.checkCacheConnection(c)
	if !ok {
		return err
	}

	ok, err = controller.checkDatabaseConnection(c, store)
	if !ok {
		return err
	}

	return c.JSON(200, &HealthCheck{
		Success: true,
		Message: fmt.Sprintf("Service %s is up and running!", module_name()),
		Version: Version,
	})
}
