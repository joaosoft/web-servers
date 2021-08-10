package routes

import (
	"web-servers/fiber/controllers"
	"web-servers/fiber/middlewares"

	"github.com/gofiber/fiber"
)

func Init(router *fiber.App) {
	router.Use(middlewares.CheckExample)

	router.Get("/v1/persons/:id_person", controllers.GetPersonByID)
	router.Get("/v1/persons/:id_person/addresses/:id_address", controllers.GetPersonAddressByID)
	router.Get("/v1/errors", controllers.GetErrorByID)
}
