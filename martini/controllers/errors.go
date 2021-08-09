package controllers

import (
	"net/http"
	"strconv"
	"web-servers/domain/models"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func GetErrorByID(req *http.Request, params martini.Params, r render.Render) {
	errorID, _ := strconv.Atoi(req.URL.Query().Get("id_error"))

	er, err := (&models.ErrorModel{}).GetErrorByID(errorID)
	if err != nil {
		r.JSON(http.StatusInternalServerError,
			ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
		return
	}

	r.JSON(http.StatusOK, er)
}
