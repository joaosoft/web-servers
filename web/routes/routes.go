package routes

import (
	"web-servers/web/controllers"
	"web-servers/web/middlewares"

	"github.com/joaosoft/web"
)

func Init(router *web.Server) {
	v1 := router.AddNamespace("v1", middlewares.CheckExample)

	v1.AddRoutes(
		web.NewRoute(web.MethodGet, "/persons/:id_person", controllers.GetPersonByID),
		web.NewRoute(web.MethodGet, "/persons/:id_person/addresses/:id_address", controllers.GetPersonAddressByID),
		web.NewRoute(web.MethodGet, "/errors", controllers.GetErrorByID),
	)
}
