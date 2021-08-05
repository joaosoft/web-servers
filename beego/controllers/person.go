package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/astaxie/beego"
)

type PersonControler struct {
	beego.Controller
}

func (c *PersonControler) GetPersonByID() {
	defer c.ServeJSON()

	age, _ := strconv.Atoi(c.Ctx.Request.URL.Query().Get("age"))
	request := GetPersonByIDRequest{
		IdPerson: c.Ctx.Input.Param(":id_person"),
		Age:      age,
	}

	fmt.Printf("> executing get person for id_person: %s", request.IdPerson)

	// ...

	c.Ctx.Output.SetStatus(http.StatusOK)
	c.Data["json"] = PersonResponse{
		Id:   request.IdPerson,
		Name: "João Ribeiro",
		Age:  request.Age,
	}
}
