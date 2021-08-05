package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"net/http"
)

type AddressControler struct {
	beego.Controller
}

func (c *AddressControler) GetPersonAddressByID(ctx *context.Context) {
	defer c.ServeJSON()

	request := GetPersonAddressByIDRequest{
		IdPerson:  ctx.Input.Param(":id_person"),
		IdAddress: ctx.Input.Param(":id_address"),
	}

	fmt.Printf("> executing get address for id_person: %s, id_address: %s", request.IdPerson, request.IdAddress)

	ctx.Output.SetStatus(http.StatusOK)
	c.Data["json"] = AddressResponse{
		Id:      request.IdAddress,
		Country: "Portugal",
		City:    "Porto",
		Street:  "Rua da calçada",
		Number:  7,
	}
}
