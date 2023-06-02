package controllers

import (
	"encoding/json"
	"github.com/joaosoft/web-servers/domain/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetPersonAddressByID(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	request := GetPersonAddressByIDRequest{
		IdPerson:  params.ByName("id_person"),
		IdAddress: params.ByName("id_address"),
	}

	w.Header().Set("Content-Type", "application/json")

	address, err := (&models.AddressModel{}).GetPersonAddressByID(request.IdPerson, request.IdAddress)
	if err != nil {
		bytes, _ := json.Marshal(
			ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		)

		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(bytes)
	}

	bytes, _ := json.Marshal(address)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(bytes)
}
