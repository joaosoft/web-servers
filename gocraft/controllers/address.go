package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gocraft/web"
	"net/http"
)

func GetPersonAddressByID(w web.ResponseWriter, req *web.Request) {
	request := GetPersonAddressByIDRequest{
		IdPerson:  req.PathParams["id_person"],
		IdAddress: req.PathParams["id_address"],
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
