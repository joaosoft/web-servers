package controllers

import (
	"encoding/json"
	"github.com/gocraft/web"
	"net/http"
)

func GetPersonAddressByID(w web.ResponseWriter, req *web.Request) {
	request := GetPersonAddressByIDRequest{
		IdPerson:  req.PathParams["id_person"],
		IdAddress: req.PathParams["id_address"],
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
