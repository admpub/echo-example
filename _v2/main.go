package main

import (
	"github.com/admpub/echo"
	"github.com/admpub/echo/engine/fasthttp"
	"github.com/admpub/echo/engine/standard"
	mw "github.com/admpub/echo/middleware"
)

func main() {
	engine := "fasthttp"

	e := echo.New()
	e.Use(mw.Log())
	e.Get("/", echo.HandlerFunc(func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	}))
	e.Get("/v2", echo.HandlerFunc(func(c echo.Context) error {
		return c.String(200, "Echo v2")
	}))
	e.Get("/ping", echo.HandlerFunc(func(c echo.Context) error {
		return c.String(200, "pong")
	}))

	switch engine {

	case "fasthttp":
		// FastHTTP
		e.Run(fasthttp.New(":4444"))

	default:
		// Standard
		e.Run(standard.New(":4444"))

	}

}
