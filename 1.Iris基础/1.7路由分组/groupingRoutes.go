package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
)

func hello() string {
	return "Hello!"
}

func main() {
	app := iris.Default()
	helloHandler := hero.Handler(hello)
	// v1 组
	v1 := app.Party("/v1")
	{
		v1.Post("/login", helloHandler)
		v1.Post("/submit", helloHandler)
		v1.Post("/read", helloHandler)
	}

	// v2 组
	v2 := app.Party("/v2")
	{
		v2.Post("/login", helloHandler)
		v2.Post("/submit", helloHandler)
		v2.Post("/read", helloHandler)
	}

	app.Run(iris.Addr(":8080"))
}
