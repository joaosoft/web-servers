package routes

import (
	"web-servers/httprouter/controllers"

	"github.com/julienschmidt/httprouter"
)

func Init(router *httprouter.Router) {
	router.GET("/v1/persons/:id_person", controllers.GetPersonByID)
	router.GET("/v1/persons/:id_person/addresses/:id_address", controllers.GetPersonAddressByID)
	router.GET("/v1/errors", controllers.GetErrorByID)
}
