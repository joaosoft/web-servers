package controllers

import (
	"github.com/joaosoft/web-servers/domain/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPersonAddressByID(ctx *gin.Context) {
	request := GetPersonAddressByIDRequest{
		IdPerson:  ctx.Param("id_person"),
		IdAddress: ctx.Param("id_address"),
	}

	ctx.Header("Content-Type", "application/json")

	address, err := (&models.AddressModel{}).GetPersonAddressByID(request.IdPerson, request.IdAddress)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
	}

	ctx.JSON(http.StatusOK, address)
}
