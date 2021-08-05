package routes

import (
	"web-servers/echo/controllers"
	"web-servers/echo/middlewares"

	"github.com/labstack/echo"
)

var (
	Router = echo.New()
)

func init() {
	Router.Use(middlewares.CheckExample)

	Router.GET("/v1/persons/:id_person", controllers.GetPersonByID)
	Router.GET("/v1/persons/:id_person/addresses/:id_address", controllers.GetPersonAddressByID)
	Router.GET("/v1/errors", controllers.GetErrorByID)
}
