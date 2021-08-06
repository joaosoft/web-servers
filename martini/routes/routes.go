package routes

import (
	"web-servers/martini/controllers"
	"web-servers/martini/middlewares"

	"github.com/martini-contrib/render"

	"github.com/go-martini/martini"
)

var (
	Router = new()
)

func new() *martini.ClassicMartini {
	r := martini.NewRouter()
	m := martini.New()
	m.Use(martini.Recovery())
	m.Use(martini.Static("public"))
	m.MapTo(r, (*martini.Routes)(nil))
	m.Action(r.Handle)
	return &martini.ClassicMartini{Martini: m, Router: r}
}

func init() {
	Router.Use(render.Renderer())
	Router.Use(middlewares.CheckExample)

	Router.Get("/v1/persons/:id_person", controllers.GetPersonByID)
	Router.Get("/v1/persons/:id_person/addresses/:id_address", controllers.GetPersonAddressByID)
	Router.Get("/v1/errors", controllers.GetErrorByID)
}
