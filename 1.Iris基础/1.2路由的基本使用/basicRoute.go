package main

import "github.com/kataras/iris"

func main() {
	app := iris.New()
	// get '/userId/12'
	app.Get("/userId/{id:uint64}", func(ctx iris.Context) {
		id := ctx.Params().GetUint64Default("id", 0) // 获取参数中类型为 uint64，字段名为 id 的值，如果没有，则返回默认值 0.
		ctx.Writef("id 是 %d", id)
	})
	// get '/userName/gbs'
	app.Get("/userName/{name}", func(ctx iris.Context) {
		name := ctx.Params().Get("name") // 获取参数中字段名为 name 的值
		ctx.Writef("name 是 %s", name)
	})
	//可以将 ctx.Params() 看作是一个 Map ，能 Get 到其中的值

	app.Get("/profile/{name:alphabetical max(255)}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
		// 规定 name 为 alphabetical 类型，并使用 max() 方法校验其最大长度不能超过255
		ctx.Writef("name 是 %s，并且没有超过255", name)
	})

	//定义验参方法 has，校验参数是否在给定的内容之中。
	app.Macros().Get("string").RegisterFunc("has", func(validNames []string) func(string) bool {
		return func(paramValue string) bool {
			for _, validName := range validNames {
				if validName == paramValue {
					return true
				}
			}
			return false
		}
	})
	// 使用 has 校验 参数值 是否在 [kataras, gerasimos, maropoulos].kataras 是Iris的作者！
	app.Get("/static_validation/{name:string has([kataras, gerasimos, maropoulos])}", func(ctx iris.Context) {

		name := ctx.Params().Get("name")
		ctx.Writef(`Hello %s | 这个名字需要是 "kataras" or "gerasimos" or "maropoulos"中的一个，否则该请求不会被处理`, name)

	})
	// 自定义 range 方法
	app.Macros().Get("string").RegisterFunc("range", func(minLength, maxLength int) func(string) bool {

		return func(paramValue string) bool {
			return len(paramValue) >= minLength && len(paramValue) <= maxLength
		}

	})
	// 使用 range 方法，如果校验不通过 则返回 status code 为400
	app.Get("/limitchar/{name:string range(1, 200) else 400}", func(ctx iris.Context) {

		name := ctx.Params().Get("name")
		ctx.Writef(`Hello %s | 这个名字长度需要为1-200之间，否则该请求不会被处理，返回400`, name)

	})

	app.Run(iris.Addr(":8080"))
}
