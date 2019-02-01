package actions

import (
	"github.com/I1820/wf/config"
	"github.com/gin-gonic/gin"
)

// App creates new version of gin router
func App() *gin.Engine {
	if config.GetConfig().Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// Default With the Logger and Recovery middleware already attached
	app := gin.Default()
	app.Use(gin.ErrorLogger())

	// content-type middleware
	app.Use(func(c *gin.Context) {
		c.Header("Content-Type", c.NegotiateFormat("application/json"))
	})

	// Routes
	app.GET("/about", AboutHandler)
	api := app.Group("/api")
	{
		api.POST("/darksky", DarkskyHandler)
	}

	return app
}
