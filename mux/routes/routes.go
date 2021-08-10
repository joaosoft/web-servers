package routes

import (
	"net/http"
	"web-servers/mux/controllers"
	"web-servers/mux/middlewares"

	"github.com/gorilla/mux"
)

func Init(router *mux.Router) {
	v1 := router.PathPrefix("/v1").Subrouter()
	v1.Use(middlewares.CheckExample)

	v1.HandleFunc("/persons/{id_person}", controllers.GetPersonByID).Methods(http.MethodGet)
	v1.HandleFunc("/persons/{id_person}/addresses/{id_address}", controllers.GetPersonAddressByID).Methods(http.MethodGet)
	v1.HandleFunc("/errors", controllers.GetErrorByID).Methods(http.MethodGet)
}
