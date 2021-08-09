package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"web-servers/domain/models"

	"github.com/gofiber/fiber/v2"
)

func GetErrorByID(ctx *fiber.Ctx) error {
	errorID, _ := strconv.Atoi(ctx.Query("id_error"))

	ctx.Response().Header.SetContentType("application/json")

	er, err := (&models.ErrorModel{}).GetErrorByID(errorID)
	if err != nil {
		ctx.Response().SetStatusCode(http.StatusInternalServerError)
		return ctx.JSON(
			ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
	}

	bytes, _ := json.Marshal(er)

	ctx.Response().SetStatusCode(http.StatusOK)
	_, err = ctx.Write(bytes)
	return err
}
