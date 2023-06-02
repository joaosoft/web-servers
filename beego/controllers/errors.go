package controllers

import (
	"github.com/joaosoft/web-servers/domain/models"
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

	c.Ctx.Output.SetStatus(http.StatusOK)
	er, err := (&models.ErrorModel{}).GetErrorByID(errorID)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	c.Data["json"] = er
}
