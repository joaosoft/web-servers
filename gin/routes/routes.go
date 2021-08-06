package routes

import (
	"net/http"
	"web-servers/gin/controllers"
	"web-servers/gin/middlewares"

	"github.com/gin-gonic/gin"
)

var (
	Router = new()
)

func new() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Recovery())
	return engine
}

func init() {
	Router.Use(middlewares.CheckExample)

	Router.Handle(http.MethodGet, "/v1/persons/:id_person", controllers.GetPersonByID)
	Router.Handle(http.MethodGet, "/v1/persons/:id_person/addresses/:id_address", controllers.GetPersonAddressByID)
	Router.Handle(http.MethodGet, "/v1/errors", controllers.GetErrorByID)
}
