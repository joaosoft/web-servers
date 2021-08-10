package routes

import (
	"web-servers/web/controllers"
	"web-servers/web/middlewares"

	"github.com/joaosoft/web"
)

func Init(router *web.Server) {
	router.AddMiddlewares(middlewares.CheckExample)

	router.AddRoutes(
		web.NewRoute(web.MethodGet, "/v1/persons/:id_person", controllers.GetPersonByID),
		web.NewRoute(web.MethodGet, "/v1/persons/:id_person/addresses/:id_address", controllers.GetPersonAddressByID),
		web.NewRoute(web.MethodGet, "/v1/errors", controllers.GetErrorByID),
	)
}
