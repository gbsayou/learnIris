package main

import "github.com/kataras/iris"

func main() {
	app := iris.Default()

	app.Get("/user", func(ctx iris.Context) {
		name := ctx.URLParamDefault("name", "Guest") //带默认值
		age := ctx.URLParam("age")                   //不带默认值
		// shortcut for ctx.Request().URL.Query().Get("lastname").
		ctx.Writef("Hello %s %s", name, age)
	})

	app.Run(iris.Addr(":8080"))
}
