package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"web-servers/domain/models"

	"github.com/julienschmidt/httprouter"
)

func GetPersonByID(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	age, _ := strconv.Atoi(req.URL.Query().Get("age"))
	request := GetPersonByIDRequest{
		IdPerson: params.ByName("id_person"),
		Age:      age,
	}

	w.Header().Set("Content-Type", "application/json")

	person, err := (&models.PersonModel{}).GetPersonByID(request.IdPerson, request.Age)
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

	bytes, _ := json.Marshal(person)
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}
