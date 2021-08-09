package routes

import (
	"web-servers/martini/controllers"
	"web-servers/martini/middlewares"

	"github.com/go-martini/martini"

	"github.com/martini-contrib/render"
)

func Init(server *martini.Martini, router martini.Router) {
	server.Use(render.Renderer())
	server.Use(middlewares.CheckExample)

	router.Get("/v1/persons/:id_person", controllers.GetPersonByID)
	router.Get("/v1/persons/:id_person/addresses/:id_address", controllers.GetPersonAddressByID)
	router.Get("/v1/errors", controllers.GetErrorByID)
}
