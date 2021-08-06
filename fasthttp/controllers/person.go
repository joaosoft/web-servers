package controllers

import (
	"encoding/json"
	routing "github.com/qiangxue/fasthttp-routing"
	"net/http"
)

func GetPersonByID(ctx *routing.Context) error {
	request := GetPersonByIDRequest{
		IdPerson: ctx.Param("id_person"),
		Age:      ctx.QueryArgs().GetUintOrZero("age"),
	}

	// ...

	ctx.SetContentType("application/json")

	bytes, err := json.Marshal(
		PersonResponse{
			Id:   request.IdPerson,
			Name: "João Ribeiro",
			Age:  request.Age,
		},
	)

	if err != nil {
		ctx.Error(err.Error(), http.StatusInternalServerError)
		return nil
	}

	ctx.SetStatusCode(http.StatusOK)
	ctx.Write(bytes)

	return nil
}
