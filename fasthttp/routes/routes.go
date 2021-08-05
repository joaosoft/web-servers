package routes

import (
	"web-servers/fasthttp/controllers"
	"web-servers/fasthttp/middlewares"

	routing "github.com/qiangxue/fasthttp-routing"
)

var (
	Router = routing.New()
)

func init() {
	Router.Use(middlewares.CheckExample)

	Router.Get("/v1/persons/<id_person>", controllers.GetPersonByID)
	Router.Get("/v1/persons/<id_person>/addresses/<id_address>", controllers.GetPersonAddressByID)
	Router.Get("/v1/errors", controllers.GetErrorByID)
}
