package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"net/http"
)

type AddressControler struct {
	beego.Controller
}

func (c *AddressControler) GetPersonAddressByID() {
	defer c.ServeJSON()

	request := GetPersonAddressByIDRequest{
		IdPerson:  c.Ctx.Input.Param(":id_person"),
		IdAddress: c.Ctx.Input.Param(":id_address"),
	}

	fmt.Printf("> executing get address for id_person: %s, id_address: %s", request.IdPerson, request.IdAddress)

	c.Ctx.Output.SetStatus(http.StatusOK)
	c.Data["json"] = AddressResponse{
		Id:      request.IdAddress,
		Country: "Portugal",
		City:    "Porto",
		Street:  "Rua da calçada",
		Number:  7,
	}
}
