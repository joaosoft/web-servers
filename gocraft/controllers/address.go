package controllers

import (
	"encoding/json"
	"github.com/joaosoft/web-servers/domain/models"
	"net/http"

	"github.com/gocraft/web"
)

func GetPersonAddressByID(w web.ResponseWriter, req *web.Request) {
	request := GetPersonAddressByIDRequest{
		IdPerson:  req.PathParams["id_person"],
		IdAddress: req.PathParams["id_address"],
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
