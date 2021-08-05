package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type ErrorControler struct {
	beego.Controller
}

func (c *ErrorControler) GetErrorByID(ctx *context.Context) {
	defer c.ServeJSON()

	errorID, _ := strconv.Atoi(ctx.Request.URL.Query().Get("id_error"))
	fmt.Printf("> executing get errors for id: %d", errorID)

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
