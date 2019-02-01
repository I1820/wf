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
	"github.com/gin-gonic/gin"
)

// darksky request payload
type darkskyReq struct {
	Latitude  float64 `json:"lat" binding:"required,min=-90,max=90"`
	Longitude float64 `json:"lng" binding:"required,min=-180,max=180"`
}

// DarkskyHandler uses darksky API to forecast weather in given geolocation.
func DarkskyHandler(c *gin.Context) {
	var rq darkskyReq

	if err := c.ShouldBind(&rq); err != nil {
		c.AbortWithError(http.StatusBadRequest, err).JSON()
		return
	}

	wr, err := darksky.ForecastRequest(rq.Latitude, rq.Longitude)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err).JSON()
		return
	}

	c.JSON(http.StatusOK, wr)
}
