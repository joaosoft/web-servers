package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gocraft/web"
)

func GetErrorByID(w web.ResponseWriter, req *web.Request) {
	errorID, _ := strconv.Atoi(req.URL.Query().Get("id_error"))
	statusText := http.StatusText(errorID)

	if statusText != "" {
		w.Header().Set("Content-Type", "application/json")
		bytes, _ := json.Marshal(
			ErrorResponse{
				Code:    errorID,
				Message: statusText,
			},
		)
		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}
