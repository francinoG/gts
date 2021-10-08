package controllers

import (
	"gts/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AddGraduates(c echo.Context) error {
	graduate := models.Graduates{}
	c.Bind(&graduate)
	err := graduate.Add()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"code":    201,
		"message": "Success",
	})
}

func GetGraduate(c echo.Context) error {
	name := c.Param("name")
	graduate := models.Graduates{}
	c.Bind(&graduate)
	grad, err := graduate.GetByName(name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    500,
			"message": err.Error(),
			"data":    "",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"message": "Success",
		"data":    grad,
	})
}
