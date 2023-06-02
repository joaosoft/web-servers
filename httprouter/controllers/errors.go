package controllers

import (
	"encoding/json"
	"github.com/joaosoft/web-servers/domain/models"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func GetErrorByID(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	errorID, _ := strconv.Atoi(req.URL.Query().Get("id_error"))

	er, err := (&models.ErrorModel{}).GetErrorByID(errorID)
	if err != nil {
		bytes, _ := json.Marshal(
			ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		)

		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(bytes)
	}

	bytes, _ := json.Marshal(er)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(bytes)
}
