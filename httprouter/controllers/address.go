package controllers

import (
	"encoding/json"
	"net/http"
	"web-servers/domain/models"

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
		w.Write(bytes)
	}

	bytes, _ := json.Marshal(address)
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}
