package controllers

import (
	"net/http"
	"strconv"
	"web-servers/implementation/models"

	"github.com/revel/revel"
)

type PersonController struct {
	*revel.Controller
}

func (c PersonController) GetPersonByID() revel.Result {
	age, _ := strconv.Atoi(c.Request.URL.Query().Get("age"))
	request := GetPersonByIDRequest{
		IdPerson: c.Params.Get("id_person"),
		Age:      age,
	}

	c.Response.WriteHeader(http.StatusOK, "application/json")

	person, err := (&models.PersonModel{}).GetPersonByID(request.IdPerson, request.Age)
	if err != nil {
		c.RenderJSON(
			ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		)
	}

	return c.RenderJSON(person)
}
