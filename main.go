package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	conf "github.com/maabodi/goApp/config"
	rest "github.com/maabodi/goApp/handler/http"

	_ "github.com/maabodi/goApp/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	config := conf.InitConfiguration()
	e := echo.New()
	e.GET("/swagger/*", echoSwagger.EchoWrapHandler())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "untuk akses swaggernya bisa diakses dibawah ini. \n\n '.../swagger/index.html#/auth/login'")
	})
	rest.RegisterUserAPI(e, config)
	e.Logger.Fatal(e.Start(":1323"))
}
