package controllers

import (
	"net/http"
	"strconv"
	"web-servers/implementation/models"

	"github.com/gin-gonic/gin"
)

func GetErrorByID(ctx *gin.Context) {
	errorID, _ := strconv.Atoi(ctx.Request.URL.Query().Get("id_error"))

	ctx.Header("Content-Type", "application/json")

	er, err := (&models.ErrorModel{}).GetErrorByID(errorID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
	}

	ctx.JSON(http.StatusOK, er)
}
