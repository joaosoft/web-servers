package controllers

import (
	"github.com/joaosoft/web-servers/domain/models"
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

	c.Ctx.Output.SetStatus(http.StatusOK)
	person, err := (&models.PersonModel{}).GetPersonByID(request.IdPerson, age)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	c.Data["json"] = person
}
