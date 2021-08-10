package routes

import (
	"web-servers/echo/controllers"
	"web-servers/echo/middlewares"

	"github.com/labstack/echo"
)

func Init(router *echo.Echo) {
	router.Use(middlewares.CheckExample)

	router.GET("/v1/persons/:id_person", controllers.GetPersonByID)
	router.GET("/v1/persons/:id_person/addresses/:id_address", controllers.GetPersonAddressByID)
	router.GET("/v1/errors", controllers.GetErrorByID)
}
