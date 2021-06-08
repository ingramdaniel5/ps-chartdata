package service

import (
	"net/http"
	"ps-chartdata/model"

	"github.com/labstack/echo"
)

// todo: refactor, this is all messed up
// GET/config handler
func SearchSymbolsHandler(c echo.Context) error {

	// todo: revise?
	search := model.Symbol{
		Name:                 "AAPL",
		ExchangeTraded:       "AAPL",
		ExchangeListed:       "AAPL",
		Timezone:             "America/New_York",
		Minmov:               1,
		Minmov2:              0,
		PointValue:           1,
		Session:              "0930-1630",
		HasIntraday:          false,
		HasNoVolume:          false,
		Description:          "test",
		Type:                 "stock",
		SupportedResolutions: []string{"1D", "2D", "3D", "W", "3W", "M", "6M"},
	}

	return c.JSON(http.StatusOK, search)
}
