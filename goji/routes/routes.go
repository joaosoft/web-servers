package routes

import (
	"github.com/joaosoft/web-servers/goji/controllers"
	"github.com/joaosoft/web-servers/goji/middlewares"

	goji "goji.io"

	"goji.io/pat"
)

func Init(router *goji.Mux) {
	router.Use(middlewares.CheckExample)

	router.HandleFunc(pat.Get("/v1/persons/:id_person"), controllers.GetPersonByID)
	router.HandleFunc(pat.Get("/v1/persons/:id_person/addresses/:id_address"), controllers.GetPersonAddressByID)
	router.HandleFunc(pat.Get("/v1/errors"), controllers.GetErrorByID)
}
