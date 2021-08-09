package controllers

import (
	"net/http"
	"web-servers/domain/models"

	"github.com/astaxie/beego"
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
	address, err := (&models.AddressModel{}).GetPersonAddressByID(request.IdPerson, request.IdAddress)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	c.Data["json"] = address
}
