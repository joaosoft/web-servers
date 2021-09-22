package routes

import (
	"net/http"
	"web-servers/gin/controllers"
	"web-servers/gin/middlewares"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	v1 := router.Group("/v1")
	v1.Use(middlewares.CheckExample)

	v1.Handle(http.MethodGet, "/persons/:id_person", controllers.GetPersonByID)
	v1.Handle(http.MethodGet, "/persons/:id_person/addresses/:id_address", controllers.GetPersonAddressByID)
	v1.Handle(http.MethodGet, "/errors", controllers.GetErrorByID)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, struct {
			Code  int    `json:"code"`
			Error string `json:"error"`
		}{
			Code:  http.StatusNotFound,
			Error: http.StatusText(http.StatusNotFound),
		})
	})
}
