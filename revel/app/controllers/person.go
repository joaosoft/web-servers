package controllers

import (
	"fmt"
	"net/http"
	"strconv"

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

	// ...

	c.Response.WriteHeader(http.StatusOK, "application/json")
	return c.RenderJSON(PersonResponse{
		Id:   request.IdPerson,
		Name: "João Ribeiro",
		Age:  request.Age,
	})
}
