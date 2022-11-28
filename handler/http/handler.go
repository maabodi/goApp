package http

import (
	"github.com/maabodi/goApp/config"
	"github.com/maabodi/goApp/database"
	co "github.com/maabodi/goApp/handler/http/controller"
	"github.com/maabodi/goApp/repository"
	"github.com/maabodi/goApp/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterUserAPI(e *echo.Echo, conf config.Config) {

	db := database.InitDB(conf)
	repo := repository.NewUserRepository(db)

	svc := service.NewServiceUser(repo, conf)

	cont := co.UserController{
		SVC: svc,
	}
	apiUser := e.Group("/api",
		middleware.Logger(),
		middleware.CORS(),
	)
	// Users Handler
	apiUser.GET("/user/all", cont.GetUsersController)
	apiUser.POST("/user/create", cont.CreateUserController)

	// Auth handler
	apiUser.POST("/login", cont.LoginUserController)

	// Admin route
	apiUser.GET("/admin/page", cont.AdminRoute, middleware.JWT([]byte(conf.JWT_KEY)))

	// User Route
}
