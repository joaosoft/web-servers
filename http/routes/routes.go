package routes

import (
	"net/http"
	"web-servers/http/controllers"
	"web-servers/http/middlewares"

	"github.com/gorilla/mux"
)

var (
	Router = mux.NewRouter()
)

func init() {
	Router.Use(middlewares.CheckExample)

	Router.HandleFunc("/v1/persons/{id_person}", controllers.GetPersonByID).Methods(http.MethodGet)
	Router.HandleFunc("/v1/persons/{id_person}/addresses/{id_address}", controllers.GetPersonAddressByID).Methods(http.MethodGet)
	Router.HandleFunc("/v1/errors", controllers.GetErrorByID).Methods(http.MethodGet)
}
