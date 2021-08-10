package routes

import (
	"web-servers/gocraft/controllers"
	"web-servers/gocraft/middlewares"

	"github.com/gocraft/web"
)

func Init(router *web.Router) {
	router.
		Middleware(middlewares.CheckExample).
		Middleware(web.ShowErrorsMiddleware).
		Get("/v1/persons/:id_person", controllers.GetPersonByID).
		Get("/v1/persons/:id_person/addresses/:id_address", controllers.GetPersonAddressByID).
		Get("/v1/errors", controllers.GetErrorByID)
}
