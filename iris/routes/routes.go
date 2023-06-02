package routes

import (
	"github.com/joaosoft/web-servers/iris/controllers"
	"github.com/joaosoft/web-servers/iris/middlewares"
)

func Init(router *iris.Application) {
	v1 := router.Party("/v1")
	v1.Use(middlewares.CheckExample)

	v1.Get("/persons/{id_person}", controllers.GetPersonByID)
	v1.Get("/persons/{id_person}/addresses/{id_address}", controllers.GetPersonAddressByID)
	v1.Get("/errors", controllers.GetErrorByID)
}
