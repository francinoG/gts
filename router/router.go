package router

import (
	"gts/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	e.POST("/graduates", controllers.AddGraduates)
	e.GET("/graduates/:name", controllers.GetGraduate)

	return e
}
