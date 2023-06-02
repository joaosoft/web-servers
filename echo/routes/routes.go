package routes

import (
	"github.com/joaosoft/web-servers/echo/controllers"
	"github.com/joaosoft/web-servers/echo/middlewares"

	"github.com/labstack/echo"
)

func Init(router *echo.Echo) {
	v1 := router.Group("/v1")
	v1.Use(middlewares.CheckExample)

	v1.GET("/persons/:id_person", controllers.GetPersonByID)
	v1.GET("/persons/:id_person/addresses/:id_address", controllers.GetPersonAddressByID)
	v1.GET("/errors", controllers.GetErrorByID)
}
