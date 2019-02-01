package actions

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AboutHandler is a default handler to serve up
// memory of 18.20.
func AboutHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "18.20 is leaving us")
}
