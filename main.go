package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	conf "github.com/myApp/config"
	rest "github.com/myApp/handler/http"
)

func main() {
	config := conf.InitConfiguration()
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hey Dunia.")
	})
	rest.RegisterUserAPI(e, config)
	e.Logger.Fatal(e.Start(":1323"))
}
