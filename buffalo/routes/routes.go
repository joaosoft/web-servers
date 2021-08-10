package routes

import (
	"web-servers/buffalo/controllers"
	"web-servers/buffalo/middlewares"

	"github.com/gobuffalo/buffalo"
)

func Init(router *buffalo.App) {
	router.Use(middlewares.CheckExample)

	router.GET("/v1/persons/{id_person}", controllers.GetPersonByID)
	router.GET("/v1/persons/{id_person}/addresses/{id_address}", controllers.GetPersonAddressByID)
	router.GET("/v1/errors", controllers.GetErrorByID)
}
