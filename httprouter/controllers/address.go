package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func GetPersonAddressByID(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	request := GetPersonAddressByIDRequest{
		IdPerson:  params.ByName("id_person"),
		IdAddress: params.ByName("id_address"),
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
