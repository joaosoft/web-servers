package routes

import (
	"web-servers/beego/controllers"
	"web-servers/beego/middlewares"

	"github.com/astaxie/beego"
)

var (
	Router = beego.BeeApp
)

func init() {
	Router.Handlers.InsertFilter("*", beego.BeforeExec, middlewares.CheckExample)

	Router.Handlers.Add("/v1/persons/:id_person", &controllers.PersonController{}, "get:GetPersonByID")
	Router.Handlers.Add("/v1/persons/:id_person/addresses/:id_address", &controllers.AddressController{}, "get:GetPersonAddressByID")
	Router.Handlers.Add("/v1/errors", &controllers.ErrorController{}, "get:GetErrorByID")
}
