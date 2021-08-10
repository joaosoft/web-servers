package routes

import (
	"web-servers/iris/controllers"
	"web-servers/iris/middlewares"

	"github.com/kataras/iris/v12"
)

func Init(router *iris.Application) {
	v1 := router.Subdomain("/v1")
	v1.Use(middlewares.CheckExample)

	v1.Get("/persons/{id_person}", controllers.GetPersonByID)
	v1.Get("/persons/{id_person}/addresses/{id_address}", controllers.GetPersonAddressByID)
	v1.Get("/errors", controllers.GetErrorByID)
}
