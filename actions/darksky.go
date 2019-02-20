/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 12-11-2018
 * |
 * | File Name:     darksky.go
 * +===============================================
 */

package actions

import (
	"net/http"

	"github.com/I1820/wf/darksky"
	"github.com/labstack/echo/v4"
)

// darksky request payload
type darkskyReq struct {
	Latitude  float64 `json:"lat" validate:"required,min=-90,max=90"`
	Longitude float64 `json:"lng" validate:"required,min=-180,max=180"`
}

// DarkskyHandler uses darksky API to forecast weather in given geolocation.
func DarkskyHandler(c echo.Context) error {
	var rq darkskyReq

	if err := c.Bind(&rq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(rq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	wr, err := darksky.ForecastRequest(rq.Latitude, rq.Longitude)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, wr)
}
