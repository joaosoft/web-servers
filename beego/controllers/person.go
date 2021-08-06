package controllers

import (
	"net/http"
	"strconv"

	"github.com/astaxie/beego"
)

type PersonController struct {
	beego.Controller
}

func (c *PersonController) GetPersonByID() {
	defer c.ServeJSON()

	age, _ := strconv.Atoi(c.Ctx.Request.URL.Query().Get("age"))
	request := GetPersonByIDRequest{
		IdPerson: c.Ctx.Input.Param(":id_person"),
		Age:      age,
	}

	// ...

	c.Ctx.Output.SetStatus(http.StatusOK)
	c.Data["json"] = PersonResponse{
		Id:   request.IdPerson,
		Name: "João Ribeiro",
		Age:  request.Age,
	}
}
