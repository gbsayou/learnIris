package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
)

//一个简单的 hello 方法，将会被 hero 转化
func hello(to string) string {
	return "Hello " + to
}

// 定义一个 service 接口，该接口需要实现一个方法 SayHello
type Service interface {
	SayHello(to string) string
}

// 定义一个结构体： myTestService，有一个属性：prefix
type myTestService struct {
	prefix string
}

// 结构体 myTestService 实现 接口中的方法
func (s *myTestService) SayHello(to string) string {
	return s.prefix + " " + to
}

// 定义一个 helloservice 方法，接收两个参数，其中一个就是那个接口。这个方法将会被 hero 转化
func helloService(to string, service Service) string {
	return service.SayHello(to)
}

//定义一个 结构体，form
type LoginForm struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

// 定义一个登录方法，将会被 hero 转化
func login(form LoginForm) string {
	return "Hello " + form.Username
}
func main() {

	app := iris.New()

	// 1.直接把hello函数转化成iris请求处理函数
	helloHandler := hero.Handler(hello)
	app.Get("/{to:string}", helloHandler)

	// 2.把结构体实例注入hero，再把结构体方法转化成iris请求处理函数
	hero.Register(&myTestService{
		prefix: "Service: Hello",
	})
	helloServiceHandler := hero.Handler(helloService)
	app.Get("/service/{to:string}", helloServiceHandler)

	// 3.注册一个iris请求处理函数，是以from表单格式x-www-form-urlencoded数据类型,以LoginForm类型映射
	// 然后把login方法转化成iris请求处理函数
	hero.Register(func(ctx iris.Context) (form LoginForm) {
		//绑定from方式提交以x-www-form-urlencoded数据格式传输的from数据，并返回相应结构体
		ctx.ReadForm(&form)
		return
	})
	loginHandler := hero.Handler(login)
	app.Post("/login", loginHandler)
	// http://localhost:8080/your_name
	// http://localhost:8080/service/your_name
	app.Run(iris.Addr(":8080"))

}
