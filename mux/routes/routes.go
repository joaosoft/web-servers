package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"web-servers/mux/controllers"
	"web-servers/mux/middlewares"
)

var (
	Router = mux.NewRouter()
)

func init() {
	Router.Use(middlewares.CheckExample)

	Router.HandleFunc("/v1/persons/{id_person}", controllers.GetPersonByID).Methods(http.MethodGet)
	Router.HandleFunc("/v1/persons/{id_person}/addresses/{id_address}", controllers.GetPersonAddressByID).Methods(http.MethodGet)
	Router.HandleFunc("/v1/errors", controllers.GetErrorByID)
}
