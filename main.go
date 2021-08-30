package main

import (
	"SC/auth"
	"SC/config"
	"SC/route"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	config.Init_DB()
	config.InitPort()
	auth.LogMiddlewares((e))
	route.New(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.HTTP_PORT)))
}
