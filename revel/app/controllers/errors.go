package controllers

import (
	"net/http"
	"strconv"
	"web-servers/domain/models"

	"github.com/revel/revel"
)

type ErrorController struct {
	*revel.Controller
}

func (c ErrorController) GetErrorByID() revel.Result {
	errorID, _ := strconv.Atoi(c.Request.URL.Query().Get("id_error"))

	c.Response.WriteHeader(http.StatusOK, "application/json")

	er, err := (&models.ErrorModel{}).GetErrorByID(errorID)
	if err != nil {
		c.RenderJSON(
			ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		)
	}

	return c.RenderJSON(er)
}
