package models

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/yashwantsinghcode/go_backend/constants"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (c *CustomValidator) Validate(i interface{}) error {
	if err := c.Validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, ErrorResponse{
			Code:    constants.ERROR_400_BAD_REQUEST,
			Message: err.Error(),
		})
	}

	return nil
}
