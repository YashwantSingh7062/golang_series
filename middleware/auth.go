package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yashwantsinghcode/go_backend/constants"
	"github.com/yashwantsinghcode/go_backend/models"
	"github.com/yashwantsinghcode/go_backend/utils"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, models.ErrorResponse{
				Code:    constants.ERROR_401_UNAUTHORIZED,
				Message: constants.ERROR_401_UNAUTHORIZED_MESSAGE,
			})
		}

		if len(tokenString) > len("Bearer ") {
			tokenString = tokenString[len("Bearer "):]
			// Add claims to the context
			_, err := utils.VerifyJwt(tokenString)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, models.ErrorResponse{
					Code:    constants.ERROR_401_UNAUTHORIZED,
					Message: constants.ERROR_401_UNAUTHORIZED_MESSAGE,
				})
			}
		} else {
			return echo.NewHTTPError(http.StatusUnauthorized, models.ErrorResponse{
				Code:    constants.ERROR_401_UNAUTHORIZED,
				Message: constants.ERROR_401_UNAUTHORIZED_MESSAGE,
			})
		}

		return next(c)
	}
}
