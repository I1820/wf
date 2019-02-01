package actions

import (
	"github.com/I1820/wf/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"gopkg.in/go-playground/validator.v9"
)

type DefaultValidator struct {
	validator *validator.Validate
}

func (cv *DefaultValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// App creates new instance of Echo and configures it
func App() *echo.Echo {
	app := echo.New()
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	if config.GetConfig().Debug {
		app.Logger.SetLevel(log.DEBUG)
	}

	// Validator
	app.Validator = &DefaultValidator{validator.New()}

	// Routes
	app.GET("/about", AboutHandler)
	api := app.Group("/api")
	{
		api.POST("/darksky", DarkskyHandler)
	}

	return app
}
