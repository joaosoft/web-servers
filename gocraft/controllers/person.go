package controllers

import (
	"encoding/json"
	"github.com/gocraft/web"
	"net/http"
	"strconv"
)

func GetPersonByID(w web.ResponseWriter, req *web.Request) {
	age, _ := strconv.Atoi(req.URL.Query().Get("age"))
	request := GetPersonByIDRequest{
		IdPerson: req.PathParams["id_person"],
		Age:      age,
	}

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
