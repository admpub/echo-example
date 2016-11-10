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

	// 静态文件服务 (URL: http://localhost:4444/static/)
	e.Use(mw.Static(&mw.StaticOptions{
		Root:"static", //存放静态文件的物理路径
		Path:"/static/", //网址访问静态文件的路径
		Browse:true, //是否在首页显示文件列表
	}))

	// ==========================
	// 设置路由
	// ==========================

	// 首页 (URL: http://localhost:4444/)
	e.Get("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!\n"+fmt.Sprintf("%+v", c.Request().Form().All()))
	})
	e.Post("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!\n"+fmt.Sprintf("%+v", c.Request().Form().All()))
	})

	// Bind (URL: http://localhost:4444/bind)
	e.Post("/bind", func(c echo.Context) error {
		m := &FormData{}
		c.Bind(m)
		return c.String(200, "Bind data:\n"+fmt.Sprintf("%+v", m))
	})

	// v2 (URL: http://localhost:4444/v2)
	e.Get("/v2", func(c echo.Context) error {
		return c.String(200, "Echo v2")
	})

	// ping (URL: http://localhost:4444/ping)
	e.Get("/ping", func(c echo.Context) error {
		return c.String(200, "pong")
	})
	
	// Stdlib handler (URL: http://localhost:4444/std)
	e.Get("/std", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`standard net/http handleFunc`))
		w.WriteHeader(200)
	})

	// ==========================
	// 创建子路由
	// ==========================

	// GET (URL: http://localhost:4444/admin)
	g := e.Group("/admin")
	g.Get("", func(c echo.Context) error {
		return c.String(200, "Hello, Group!\n"+fmt.Sprintf("%+v", c.Request().Form().All()))
	})

	// POST (URL: http://localhost:4444/admin)
	g.Post("", func(c echo.Context) error {
		return c.String(200, "Hello, Group!\n"+fmt.Sprintf("%+v", c.Request().Form().All()))
	})

	// (URL: http://localhost:4444/admin/ping)
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
