package service

import (
	"net/http"
	"ps-chartdata/model"

	"github.com/labstack/echo"
)

// GET/config handler
func GetConfigHandler(c echo.Context) error {

	// todo: revise?
	config := model.Config{
		SupportedResolutions:   []string{"D", "2D", "3D", "W", "3W", "M", "6M"},
		SupportsGroupRequest:   false,
		SupportsMarks:          false,
		SupportsSearch:         true,
		SupportsTimescaleMarks: false,
		SupportsTime:           false,
	}

	return c.JSON(http.StatusOK, config)
}
