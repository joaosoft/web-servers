package routes

import (
	"github.com/joaosoft/web-servers/mux/controllers"
	"github.com/joaosoft/web-servers/mux/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func Init(router *mux.Router) {
	v1 := router.PathPrefix("/v1").Subrouter()
	v1.Use(middlewares.CheckExample)

	v1.HandleFunc("/persons/{id_person}", controllers.GetPersonByID).Methods(http.MethodGet)
	v1.HandleFunc("/persons/{id_person}/addresses/{id_address}", controllers.GetPersonAddressByID).Methods(http.MethodGet)
	v1.HandleFunc("/errors", controllers.GetErrorByID).Methods(http.MethodGet)
}
