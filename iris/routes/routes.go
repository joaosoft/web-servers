package routes

import (
	"web-servers/iris/controllers"
	"web-servers/iris/middlewares"

	"github.com/kataras/iris"
)

func Init(router *iris.Application) {
	router.Use(middlewares.CheckExample)

	router.Get("/v1/persons/{id_person}", controllers.GetPersonByID)
	router.Get("/v1/persons/{id_person}/addresses/{id_address}", controllers.GetPersonAddressByID)
	router.Get("/v1/errors", controllers.GetErrorByID)
}
