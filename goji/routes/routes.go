package routes

import (
	"web-servers/goji/controllers"
	"web-servers/goji/middlewares"

	"goji.io"
	"goji.io/pat"
)

var (
	Router = goji.NewMux()
)

func init() {
	Router.Use(middlewares.CheckExample)

	Router.HandleFunc(pat.Get("/v1/persons/:id_person"), controllers.GetPersonByID)
	Router.HandleFunc(pat.Get("/v1/persons/:id_person/addresses/:id_address"), controllers.GetPersonAddressByID)
	Router.HandleFunc(pat.Get("/v1/errors"), controllers.GetErrorByID)
}
