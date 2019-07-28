package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func hello() string {
	return "Hello!"
}
func MyBenchLogger(ctx iris.Context) {
	fmt.Println("收到一个请求")
}

func main() {
	app := iris.New()

	helloHandler := hero.Handler(hello)
	// app := iris.Default() 则默认使用 recover 和 logger 两个中间件。本节想要自定义 logger

	// recover中间件处理所有的panic报错，并且返回一个500的 status code
	app.Use(recover.New())

	// 自定义 logger
	requestLogger := logger.New(logger.Config{
		// Status 展示 status code
		Status: true,
		// IP 请求的IP地址
		IP: true,
		// Method 请求的方法 post get 等
		Method: true,
		// Path 请求路径
		Path: true,
		// Query 将请求所带的参数拼在路径中 /user?name=gbs
		Query: true,
		// if !empty then its contents derives from `ctx.Values().Get("logger_message")
		// will be added to the logs.
		MessageContextKeys: []string{"logger_message"},

		// if !empty then its contents derives from `ctx.GetHeader("User-Agent")
		MessageHeaderKeys: []string{"User-Agent"},
	})
	// 全局使用该中间件
	app.Use(requestLogger)

	// 以下示例是非全局中间件
	// 单个路由中使用 MyBenchLogger，如果要使用多个中间件，在最后的处理方法之前添加即可

	app.Get("/benchmark", MyBenchLogger, helloHandler)

	// 分组路由使用中间件
	authorized := app.Party("/user")
	authorized.Use(MyBenchLogger)
	// 以上代码也可以写成：
	// authorized := app.Party("/user", AuthRequired())
	// exactly the same as:
	{
		authorized.Post("/login", helloHandler)
		authorized.Post("/submit", helloHandler)
		authorized.Post("/read", helloHandler)

		// 分组嵌套: /user/testing
		testing := authorized.Party("/testing")
		testing.Get("/analytics", helloHandler)
	}

	app.Run(iris.Addr(":8080"))
}
