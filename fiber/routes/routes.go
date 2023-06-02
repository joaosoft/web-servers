package routes

import (
	"github.com/joaosoft/web-servers/fiber/controllers"
	"github.com/joaosoft/web-servers/fiber/middlewares"
)

func Init(router *fiber.App) {
	v1 := router.Group("/v1")
	v1.Use(middlewares.CheckExample)

	v1.Get("/persons/:id_person", controllers.GetPersonByID)
	v1.Get("/persons/:id_person/addresses/:id_address", controllers.GetPersonAddressByID)
	v1.Get("/errors", controllers.GetErrorByID)
}
