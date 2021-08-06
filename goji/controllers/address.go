package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"goji.io/pat"
)

func GetPersonAddressByID(w http.ResponseWriter, req *http.Request) {
	request := GetPersonAddressByIDRequest{
		IdPerson:  pat.Param(req, "id_person"),
		IdAddress: pat.Param(req, "id_address"),
	}

	fmt.Printf("> executing get address for id_person: %s, id_address: %s", request.IdPerson, request.IdAddress)

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
