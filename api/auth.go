package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yashwantsinghcode/go_backend/constants"
	"github.com/yashwantsinghcode/go_backend/models"
	"github.com/yashwantsinghcode/go_backend/utils"
)

var email = "yashwant@gmail.com"
var password = "password"

func (a *Api) Login(ctx echo.Context) error {
	// Binding request body
	request := new(models.LoginRequest)
	if err := ctx.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrorResponse{
			Code:    constants.ERROR_400_BAD_REQUEST,
			Message: constants.ERROR_400_BAD_REQUEST_MESSAGE,
		})
	}

	// Validating request body
	if err := ctx.Validate(request); err != nil {
		return err
	}

	// To do: Check the user in the DB.
	// To do: Create a hashed password
	if email == request.Email && password == request.Password {
		// Token Generation
		claims := map[string]interface{}{
			"_id":   "__unique_id__",
			"email": email,
		}
		token, err := utils.SignJwt(claims)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, models.ErrorResponse{
				Code:    constants.ERROR_401_UNAUTHORIZED,
				Message: constants.ERROR_401_UNAUTHORIZED_MESSAGE,
			})
		}
		return ctx.JSON(http.StatusOK, models.LoginResponse{
			Message: "Login Successful",
			Token:   token,
		})
	} else {
		return echo.NewHTTPError(http.StatusUnauthorized, models.ErrorResponse{
			Code:    constants.ERROR_401_INVALID_CREDENTIALS,
			Message: constants.ERROR_401_INVALID_CREDENTIALS_MESSAGE,
		})
	}
}

func (a *Api) Signup(ctx echo.Context) error {
	// Binding Request
	user := new(models.User)
	if err := ctx.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrorResponse{
			Code:    constants.ERROR_400_BAD_REQUEST,
			Message: constants.ERROR_400_BAD_REQUEST_MESSAGE,
		})
	}

	// Validating Request
	if err := ctx.Validate(user); err != nil {
		return err
	}

	// Creating new User
	// To do: Check user exist in DB
	// To do: Create new user in DB
	user.Id = "New_Generated_Id"

	return ctx.JSON(http.StatusOK, models.UserResponse{
		Message: "Signed up successfully!.",
		Data:    user,
	})
}

func (a *Api) Profile(ctx echo.Context) error {
	userId := ctx.Param("id")

	// To do: Fetch User Details from the DB.
	// Verify if user exist

	return ctx.JSON(http.StatusOK, models.UserResponse{
		Message: "Profile fetched successfully!.",
		Data: &models.User{
			Id:       userId,
			Name:     "Yashwant Singh",
			Email:    email,
			Password: password,
		},
	})
}
