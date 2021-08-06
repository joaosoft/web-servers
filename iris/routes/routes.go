package routes

import (
	"web-servers/iris/controllers"
	"web-servers/iris/middlewares"

	"github.com/kataras/iris/v12"
)

var (
	Router = iris.New()
)

func init() {
	Router.Use(middlewares.CheckExample)

	Router.Get("/v1/persons/{id_person}", controllers.GetPersonByID)
	Router.Get("/v1/persons/{id_person}/addresses/{id_address}", controllers.GetPersonAddressByID)
	Router.Get("/v1/errors", controllers.GetErrorByID)
}
