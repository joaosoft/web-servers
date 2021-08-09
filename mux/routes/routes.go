package routes

import (
	"net/http"
	"web-servers/mux/controllers"
	"web-servers/mux/middlewares"

	"github.com/gorilla/mux"
)

func Init(router *mux.Router) {
	router.Use(middlewares.CheckExample)

	router.HandleFunc("/v1/persons/{id_person}", controllers.GetPersonByID).Methods(http.MethodGet)
	router.HandleFunc("/v1/persons/{id_person}/addresses/{id_address}", controllers.GetPersonAddressByID).Methods(http.MethodGet)
	router.HandleFunc("/v1/errors", controllers.GetErrorByID).Methods(http.MethodGet)
}
