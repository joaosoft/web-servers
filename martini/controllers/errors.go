package controllers

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
	"strconv"
)

func GetErrorByID(req *http.Request, params martini.Params, r render.Render) {
	errorID, _ := strconv.Atoi(req.URL.Query().Get("id_error"))
	statusText := http.StatusText(errorID)

	if statusText != "" {
		r.JSON(http.StatusOK, ErrorResponse{
			Code:    errorID,
			Message: statusText,
		})
	} else {
		r.Status(http.StatusNoContent)
	}
}
