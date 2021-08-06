package routes

import (
	"web-servers/buffalo/controllers"
	"web-servers/buffalo/middlewares"

	"github.com/gobuffalo/buffalo"
)

var (
	Router = buffalo.New(buffalo.Options{Addr: ":8081"})
)

func init() {
	Router.Use(middlewares.CheckExample)

	Router.GET("/v1/persons/{id_person}", controllers.GetPersonByID)
	Router.GET("/v1/persons/{id_person}/addresses/{id_address}", controllers.GetPersonAddressByID)
	Router.GET("/v1/errors", controllers.GetErrorByID)
}
