package controller

import (
	"fmt"
	"runtime/debug"

	"github.com/labstack/echo/v4"
)

type HealthCheckController struct{}

type HealthCheck struct {
	Success bool   `json:"success" xml:"success"`
	Message string `json:"message" xml:"message"`
	Version string `json:"version" xml:"version"`
}

var commit_hash = func() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			if setting.Key == "vcs.revision" {
				return setting.Value[:7]
			}
		}
	}
	return ""
}()

var module_name = func() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		return info.Main.Path
	}
	return ""
}

func (controller HealthCheckController) Execute(c echo.Context) error {
	return c.JSON(200, &HealthCheck{
		Success: true,
		Message: fmt.Sprintf("Service %s is up and running!", module_name()),
		Version: commit_hash,
	})
}
