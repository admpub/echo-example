package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/admpub/fasthttp"
)

func main() {
	port := flag.String("p", "5000", "port of your app.")
	flag.Parse()

	fasthttp.ListenAndServe(`127.0.0.1:`+*port, func(c *fasthttp.RequestCtx) {
		defer func() {
			c.Success(`text/html`, []byte(`Hello world.`))
		}()
		fmt.Printf("%#v\n", c.Request.Header.String())
		form, err := c.MultipartForm() //BUG
		if err != nil {
			fmt.Println(err)
		}
		if form == nil {
			return
		}
		b, err := json.MarshalIndent(form.Value, ``, `  `)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(b))
	})
}
