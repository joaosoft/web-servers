package controllers

import (
	"encoding/json"
	"net/http"

	routing "github.com/qiangxue/fasthttp-routing"
)

func GetErrorByID(ctx *routing.Context) error {
	errorID := ctx.QueryArgs().GetUintOrZero("id_error")
	statusText := http.StatusText(errorID)

	if statusText != "" {
		ctx.SetContentType("application/json")
		bytes, err := json.Marshal(
			ErrorResponse{
				Code:    errorID,
				Message: statusText,
			},
		)

		if err != nil {
			ctx.Error(err.Error(), http.StatusInternalServerError)
			return nil
		}

		ctx.SetStatusCode(http.StatusOK)
		ctx.Write(bytes)
	} else {
		ctx.SetStatusCode(http.StatusNoContent)
	}

	return nil
}
