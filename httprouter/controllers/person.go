package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func GetPersonByID(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	age, _ := strconv.Atoi(req.URL.Query().Get("age"))
	request := GetPersonByIDRequest{
		IdPerson: params.ByName("id_person"),
		Age:      age,
	}

	fmt.Printf("> executing get person for id_person: %s", request.IdPerson)

	// ...

	bytes, _ := json.Marshal(
		PersonResponse{
			Id:   request.IdPerson,
			Name: "João Ribeiro",
			Age:  request.Age,
		},
	)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}
