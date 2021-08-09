package routes

import (
	"web-servers/fasthttp/controllers"
	"web-servers/fasthttp/middlewares"

	routing "github.com/qiangxue/fasthttp-routing"
)

func Init(router *routing.Router) {
	router.Use(middlewares.CheckExample)

	router.Get("/v1/persons/<id_person>", controllers.GetPersonByID)
	router.Get("/v1/persons/<id_person>/addresses/<id_address>", controllers.GetPersonAddressByID)
	router.Get("/v1/errors", controllers.GetErrorByID)
}
