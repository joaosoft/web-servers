package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/martini-contrib/render"

	"github.com/go-martini/martini"
)

func GetPersonByID(req *http.Request, params martini.Params, r render.Render) {
	age, _ := strconv.Atoi(req.URL.Query().Get("age"))
	request := GetPersonByIDRequest{
		IdPerson: params["id_person"],
		Age:      age,
	}

	fmt.Printf("> executing get person for id_person: %s", request.IdPerson)

	// ...

	r.JSON(http.StatusOK, PersonResponse{
		Id:   request.IdPerson,
		Name: "João Ribeiro",
		Age:  request.Age,
	})
}
