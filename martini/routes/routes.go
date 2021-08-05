package routes

import (
	"web-servers/martini/controllers"
	"web-servers/martini/middlewares"

	"github.com/martini-contrib/render"

	"github.com/go-martini/martini"
)

var (
	Router = martini.Classic()
)

func init() {
	Router.Use(render.Renderer())
	Router.Use(middlewares.CheckExample)

	Router.Get("/v1/persons/:id_person", controllers.GetPersonByID)
	Router.Get("/v1/persons/:id_person/addresses/:id_address", controllers.GetPersonAddressByID)
	Router.Get("/v1/errors", controllers.GetErrorByID)
}
