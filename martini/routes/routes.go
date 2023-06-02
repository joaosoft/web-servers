package routes

import (
	"github.com/joaosoft/web-servers/martini/controllers"
	"github.com/joaosoft/web-servers/martini/middlewares"

	"github.com/go-martini/martini"

	"github.com/martini-contrib/render"
)

func Init(server *martini.Martini, router martini.Router) {
	server.Use(render.Renderer())
	server.Use(middlewares.CheckExample)

	router.Group("/v1", func(r martini.Router) {
		r.Get("/persons/:id_person", controllers.GetPersonByID)
		r.Get("/persons/:id_person/addresses/:id_address", controllers.GetPersonAddressByID)
		r.Get("/errors", controllers.GetErrorByID)
	})
}
