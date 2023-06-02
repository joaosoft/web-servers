package routes

import (
	"github.com/joaosoft/web-servers/httprouter/controllers"
	"github.com/joaosoft/web-servers/httprouter/middlewares"
)

func Init(router *middlewares.RouterWrapper) {
	router.Use(middlewares.CheckExample)

	router.GET("/v1/persons/:id_person", controllers.GetPersonByID)
	router.GET("/v1/persons/:id_person/addresses/:id_address", controllers.GetPersonAddressByID)
	router.GET("/v1/errors", controllers.GetErrorByID)
}
