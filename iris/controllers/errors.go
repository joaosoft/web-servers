package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/kataras/iris/v12"
)

func GetErrorByID(ctx iris.Context) {
	errorID, _ := strconv.Atoi(ctx.URLParam("id_error"))
	fmt.Printf("> executing get errors for id: %d", errorID)

	statusText := http.StatusText(errorID)

	if statusText != "" {
		response := ErrorResponse{
			Code:    errorID,
			Message: statusText,
		}
		ctx.StatusCode(http.StatusOK)
		ctx.JSON(response)
	} else {
		ctx.StatusCode(http.StatusNoContent)
	}
}
