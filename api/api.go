package api

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/yashwantsinghcode/go_backend/middleware"
	"github.com/yashwantsinghcode/go_backend/models"
)

type Api struct {
	Config *models.Config
	Server *echo.Echo
	//To Do Add DB Repo
}

func NewService(cfg *models.Config) *Api {
	// To Do add db repository
	serviceV1 := &Api{
		Config: cfg,
	}
	return serviceV1
}

func (a *Api) RunHttpServer() {
	a.Server = echo.New()
	// Custom Request Validator
	a.Server.Validator = &models.CustomValidator{Validator: validator.New()}
	a.InitializeRoutes()

	a.Server.Logger.Fatal(a.Server.Start(a.Config.Address))
}

func (a *Api) InitializeRoutes() {
	// Public Routes
	publicRoutes := a.Server.Group("/")
	publicRoutes.POST("login", a.Login)
	publicRoutes.POST("signup", a.Signup)

	// Private Routes
	privateRoutes := a.Server.Group("/")
	privateRoutes.Use(middleware.Auth)
	privateRoutes.GET("profile/:id", a.Profile)
}
