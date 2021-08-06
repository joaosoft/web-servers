package controllers

import (
	"github.com/astaxie/beego"
	"net/http"
)

type AddressController struct {
	beego.Controller
}

func (c *AddressController) GetPersonAddressByID() {
	defer c.ServeJSON()

	request := GetPersonAddressByIDRequest{
		IdPerson:  c.Ctx.Input.Param(":id_person"),
		IdAddress: c.Ctx.Input.Param(":id_address"),
	}

	c.Ctx.Output.SetStatus(http.StatusOK)
	c.Data["json"] = AddressResponse{
		Id:      request.IdAddress,
		Country: "Portugal",
		City:    "Porto",
		Street:  "Rua da calçada",
		Number:  7,
	}
}
