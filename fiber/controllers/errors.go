package controllers

import (
	"net/http"
	"strconv"
	"web-servers/domain/models"

	"github.com/gofiber/fiber"
)

func GetErrorByID(ctx *fiber.Ctx) {
	errorID, _ := strconv.Atoi(ctx.Query("id_error"))

	ctx.Set("Content-Type", "application/json")

	er, err := (&models.ErrorModel{}).GetErrorByID(errorID)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		_ = ctx.JSON(
			ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
		return
	}

	ctx.Status(http.StatusOK)
	_ = ctx.JSON(er)
}
