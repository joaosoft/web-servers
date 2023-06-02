package routes

import (
	"github.com/joaosoft/web-servers/gocraft/controllers"
	"github.com/joaosoft/web-servers/gocraft/middlewares"

	"github.com/gocraft/web"
)

func Init(context interface{}, router *web.Router) {
	v1 := router.Subrouter(context, "/v1")
	v1.Middleware(middlewares.CheckExample).
		Middleware(web.ShowErrorsMiddleware).
		Get("/persons/:id_person", controllers.GetPersonByID).
		Get("/persons/:id_person/addresses/:id_address", controllers.GetPersonAddressByID).
		Get("/errors", controllers.GetErrorByID)
}
