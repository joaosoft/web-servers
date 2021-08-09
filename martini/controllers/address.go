package controllers

import (
	"net/http"
	"web-servers/implementation/models"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func GetPersonAddressByID(params martini.Params, r render.Render) {
	request := GetPersonAddressByIDRequest{
		IdPerson:  params["id_person"],
		IdAddress: params["id_address"],
	}

	address, err := (&models.AddressModel{}).GetPersonAddressByID(request.IdPerson, request.IdAddress)
	if err != nil {
		r.JSON(http.StatusInternalServerError,
			ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
		return
	}

	r.JSON(http.StatusOK, address)
}
