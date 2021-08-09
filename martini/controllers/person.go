package controllers

import (
	"net/http"
	"strconv"
	"web-servers/implementation/models"

	"github.com/martini-contrib/render"

	"github.com/go-martini/martini"
)

func GetPersonByID(req *http.Request, params martini.Params, r render.Render) {
	age, _ := strconv.Atoi(req.URL.Query().Get("age"))
	request := GetPersonByIDRequest{
		IdPerson: params["id_person"],
		Age:      age,
	}

	person, err := (&models.PersonModel{}).GetPersonByID(request.IdPerson, age)
	if err != nil {
		r.JSON(http.StatusInternalServerError,
			ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
		return
	}

	r.JSON(http.StatusOK, person)
}
