package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetPersonAddressByID(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	request := GetPersonAddressByIDRequest{
		IdPerson:  vars["id_person"],
		IdAddress: vars["id_address"],
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
