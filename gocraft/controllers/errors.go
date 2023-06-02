package controllers

import (
	"encoding/json"
	"github.com/joaosoft/web-servers/domain/models"
	"net/http"
	"strconv"

	"github.com/gocraft/web"
)

func GetErrorByID(w web.ResponseWriter, req *web.Request) {
	errorID, _ := strconv.Atoi(req.URL.Query().Get("id_error"))
	w.Header().Set("Content-Type", "application/json")

	er, err := (&models.ErrorModel{}).GetErrorByID(errorID)
	if err != nil {
		bytes, _ := json.Marshal(
			ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write(bytes)
	}

	bytes, _ := json.Marshal(er)
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}
