package routes

import (
	"encoding/json"
	"net/http"
	"web-servers/beego/controllers"
	"web-servers/beego/middlewares"

	"github.com/astaxie/beego"
)

func Init(router *beego.App) {
	_ = router.Handlers.InsertFilter("*", beego.BeforeExec, middlewares.CheckExample)
	router.Handlers.Add("/v1/persons/:id_person", &controllers.PersonController{}, "get:GetPersonByID")
	router.Handlers.Add("/v1/persons/:id_person/addresses/:id_address", &controllers.AddressController{}, "get:GetPersonAddressByID")
	router.Handlers.Add("/v1/errors", &controllers.ErrorController{}, "get:GetErrorByID")

	beego.ErrorHandler("404",
		func(w http.ResponseWriter, r *http.Request) {
			bytes, _ := json.Marshal(struct {
				Code  int    `json:"code"`
				Error string `json:"error"`
			}{
				Code:  http.StatusNotFound,
				Error: http.StatusText(http.StatusNotFound),
			})
			_, _ = w.Write(bytes)
		},
	)
}
