package routes

import (
	"web-servers/buffalo/controllers"
	"web-servers/buffalo/middlewares"

	"github.com/gobuffalo/buffalo"
)

func Init(router *buffalo.App) {
	v1 := router.Group("/v1")
	v1.Use(middlewares.CheckExample)

	v1.GET("/persons/{id_person}", controllers.GetPersonByID)
	v1.GET("/persons/{id_person}/addresses/{id_address}", controllers.GetPersonAddressByID)
	v1.GET("/errors", controllers.GetErrorByID)
}
