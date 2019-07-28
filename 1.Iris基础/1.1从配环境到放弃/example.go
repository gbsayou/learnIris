package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	app.Use(recover.New()) // 从错误中恢复
	app.Use(logger.New())  // 记录日志
	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome</h1>")
	})
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed)) // 跑起来
}
