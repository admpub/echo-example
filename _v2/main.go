package main

import (
	"fmt"
	"net/http"

	"github.com/webx-top/echo"
	"github.com/webx-top/echo/engine/fasthttp"
	"github.com/webx-top/echo/engine/standard"
	mw "github.com/webx-top/echo/middleware"
)

type FormData struct {
	User  string
	Id    int
	Email string
}

func main() {
	engine := "fasthttp"

	e := echo.New()
	e.SetDebug(true)
	// ==========================
	// 添加中间件
	// ==========================
	e.Use(func(h echo.Handler) echo.HandlerFunc {
		return func(c echo.Context) error {
			fmt.Println(`==========before===========`)
			err := h.Handle(c)
			fmt.Println(`===========after===========`)
			fmt.Println(`===========response content:`)
			fmt.Println(string(c.Response().Body()))
			return err
		}
	})

	e.Use(mw.Log())
	e.Use(mw.Recover())
	e.Use(mw.Gzip())

	// ==========================
	// 设置路由
	// ==========================
	e.Get("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!\n"+fmt.Sprintf("%+v", c.Request().Form().All()))
	})
	e.Post("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!\n"+fmt.Sprintf("%+v", c.Request().Form().All()))
	})
	e.Post("/bind", func(c echo.Context) error {
		m := &FormData{}
		c.Bind(m)
		return c.String(200, "Bind data:\n"+fmt.Sprintf("%+v", m))
	})
	e.Get("/v2", func(c echo.Context) error {
		return c.String(200, "Echo v2")
	})
	e.Get("/ping", func(c echo.Context) error {
		return c.String(200, "pong")
	})
	e.Get("/std", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`standard net/http handleFunc`))
		w.WriteHeader(200)
	})

	// ==========================
	// 创建子路由
	// ==========================
	g := e.Group("/admin")
	g.Get("", func(c echo.Context) error {
		return c.String(200, "Hello, Group!\n"+fmt.Sprintf("%+v", c.Request().Form().All()))
	})
	g.Post("", func(c echo.Context) error {
		return c.String(200, "Hello, Group!\n"+fmt.Sprintf("%+v", c.Request().Form().All()))
	})
	g.Get("/ping", func(c echo.Context) error {
		return c.String(200, "pong -- Group")
	})

	// ==========================
	// 启动服务
	// ==========================
	switch engine {

	case "fasthttp":
		// FastHTTP
		e.Run(fasthttp.New(":4444"))

	default:
		// Standard
		e.Run(standard.New(":4444"))

	}

}
