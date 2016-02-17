package main

import (
	"fmt"

	"github.com/admpub/echo"
	"github.com/admpub/echo/engine/fasthttp"
	"github.com/admpub/echo/engine/standard"
	mw "github.com/admpub/echo/middleware"
)

type FormData struct {
	User  string
	Id    int
	Email string
}

func main() {
	engine := "fasthttp"

	e := echo.New()
	e.Use(echo.MiddlewareFunc(func(h echo.Handler) echo.Handler {
		return echo.HandlerFunc(func(c echo.Context) error {
			fmt.Println(`==========before===========`)
			err := h.Handle(c)
			fmt.Println(`===========after===========`)
			return err
		})
	}))
	e.Use(mw.Log())
	e.Use(mw.Recover())
	e.Use(mw.Gzip())
	e.Get("/", echo.HandlerFunc(func(c echo.Context) error {
		return c.String(200, "Hello, World!\n"+fmt.Sprintf("%+v", c.Request().Form().All()))
	}))
	e.Post("/", echo.HandlerFunc(func(c echo.Context) error {
		return c.String(200, "Hello, World!\n"+fmt.Sprintf("%+v", c.Request().Form().All()))
	}))
	e.Post("/bind", echo.HandlerFunc(func(c echo.Context) error {
		m := &FormData{}
		c.Bind(m)
		return c.String(200, "Bind data:\n"+fmt.Sprintf("%+v", m))
	}))
	e.Get("/v2", echo.HandlerFunc(func(c echo.Context) error {
		return c.String(200, "Echo v2")
	}))
	e.Get("/ping", echo.HandlerFunc(func(c echo.Context) error {
		return c.String(200, "pong")
	}))
	g := e.Group("/admin")
	g.Get("", echo.HandlerFunc(func(c echo.Context) error {
		return c.String(200, "Hello, Group!\n"+fmt.Sprintf("%+v", c.Request().Form().All()))
	}))
	g.Post("", echo.HandlerFunc(func(c echo.Context) error {
		return c.String(200, "Hello, Group!\n"+fmt.Sprintf("%+v", c.Request().Form().All()))
	}))
	g.Get("/ping", echo.HandlerFunc(func(c echo.Context) error {
		return c.String(200, "pong -- Group")
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
