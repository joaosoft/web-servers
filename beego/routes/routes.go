package routes

import (
	"web-servers/beego/controllers"
	"web-servers/beego/middlewares"

	"github.com/astaxie/beego"
)

func Init(router *beego.App) {
	_ = router.Handlers.InsertFilter("*", beego.BeforeExec, middlewares.CheckExample)
	router.Handlers.Add("/v1/persons/:id_person", &controllers.PersonController{}, "get:GetPersonByID")
	router.Handlers.Add("/v1/persons/:id_person/addresses/:id_address", &controllers.AddressController{}, "get:GetPersonAddressByID")
	router.Handlers.Add("/v1/errors", &controllers.ErrorController{}, "get:GetErrorByID")
}
