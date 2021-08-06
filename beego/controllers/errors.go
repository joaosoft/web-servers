package controllers

import (
	"net/http"
	"strconv"

	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) GetErrorByID() {
	defer c.ServeJSON()

	errorID, _ := strconv.Atoi(c.Ctx.Request.URL.Query().Get("id_error"))
	statusText := http.StatusText(errorID)

	if statusText != "" {
		c.Ctx.Output.SetStatus(http.StatusOK)
		c.Data["json"] = ErrorResponse{
			Code:    errorID,
			Message: statusText,
		}
	} else {
		c.Ctx.Output.SetStatus(http.StatusNoContent)
	}
}
