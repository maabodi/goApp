package main

import (
	"net"
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
	// e.Logger.Fatal(e.Start(":1323"))

	l, err := net.Listen("tcp", ":1323")
	if err != nil {
		e.Logger.Fatal(err)
	}
	e.Listener = l
	e.Logger.Fatal(e.Start(""))
}
