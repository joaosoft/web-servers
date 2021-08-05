package routes

import (
	"web-servers/httprouter/controllers"
	"web-servers/httprouter/middlewares"
)

var (
	Router = middlewares.NewRouter(middlewares.CheckExample)
)

func init() {
	Router.GET("/v1/persons/:id_person", controllers.GetPersonByID)
	Router.GET("/v1/persons/:id_person/addresses/:id_address", controllers.GetPersonAddressByID)
	Router.GET("/v1/errors", controllers.GetErrorByID)
}
