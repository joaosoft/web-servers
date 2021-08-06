package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetPersonAddressByID(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	request := GetPersonAddressByIDRequest{
		IdPerson:  params.ByName("id_person"),
		IdAddress: params.ByName("id_address"),
	}

	bytes, _ := json.Marshal(
		AddressResponse{
			Id:      request.IdAddress,
			Country: "Portugal",
			City:    "Porto",
			Street:  "Rua da calçada",
			Number:  7,
		},
	)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}
