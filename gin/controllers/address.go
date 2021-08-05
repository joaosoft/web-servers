package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPersonAddressByID(ctx *gin.Context) {
	request := GetPersonAddressByIDRequest{
		IdPerson:  ctx.Param(":id_person"),
		IdAddress: ctx.Param(":id_address"),
	}

	fmt.Printf("> executing get address for id_person: %s, id_address: %s", request.IdPerson, request.IdAddress)

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, AddressResponse{
		Id:      request.IdAddress,
		Country: "Portugal",
		City:    "Porto",
		Street:  "Rua da calçada",
		Number:  7,
	})
}
