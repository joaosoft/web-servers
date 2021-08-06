package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/revel/revel"
)

type ErrorController struct {
	*revel.Controller
}

func (c ErrorController) GetErrorByID() revel.Result {
	errorID, _ := strconv.Atoi(c.Request.URL.Query().Get("id_error"))
	fmt.Printf("> executing get errors for id: %d", errorID)

	statusText := http.StatusText(errorID)

	if statusText != "" {
		c.Response.WriteHeader(http.StatusOK, "application/json")
		return c.RenderJSON(ErrorResponse{
			Code:    errorID,
			Message: statusText,
		})
	} else {
		c.Response.SetStatus(http.StatusNoContent)
		return c.Result
	}
}