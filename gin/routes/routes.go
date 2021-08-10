package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web-servers/gin/controllers"
	"web-servers/gin/middlewares"
)

func Init(router *gin.Engine) {
	router.Use(middlewares.CheckExample)

	router.Handle(http.MethodGet, "/v1/persons/:id_person", controllers.GetPersonByID)
	router.Handle(http.MethodGet, "/v1/persons/:id_person/addresses/:id_address", controllers.GetPersonAddressByID)
	router.Handle(http.MethodGet, "/v1/errors", controllers.GetErrorByID)
}
