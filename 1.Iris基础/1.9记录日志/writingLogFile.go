package main

import (
	"github.com/kataras/iris"
	"io"
	"os"
	"time"

	"github.com/kataras/iris/middleware/logger"
)

// 日志文件的名称。
func todayFilename() string {
	today := time.Now().Format("Jan 02 2006")
	return today + ".txt"
}

// 生成记录日志的文件
func newLogFile() *os.File {
	filename := todayFilename()
	// 打开该文件，如果不存在 则新建，如果已经存在，则将日志追加入其中
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	return f
}
func main() {
	f := newLogFile()
	defer f.Close()
	app := iris.New()
	requestLogger := logger.New(logger.Config{
		Status: true,
		IP:     true,
		Method: true,
		Path:   true,
		Query:  true,
		// if !empty then its contents derives from `ctx.Values().Get("logger_message")
		// will be added to the logs.
		MessageContextKeys: []string{"logger_message"},
		// if !empty then its contents derives from `ctx.GetHeader("User-Agent")
		MessageHeaderKeys: []string{"User-Agent"},
	})
	app.Use(requestLogger) //配置了log的格式之后 会主动将这些内容写入log

	// 将文件附加为logger的输出目标，iris 的logger就是一个io.Writer，可以将内容输出到文件中或者控制台
	// 将日志同时输出到控制台和日志文件中
	app.Logger().SetOutput(io.MultiWriter(f, os.Stdout))
	// app.Logger().SetOutput(f) //仅将日志输出到文件中

	app.Get("/ping", func(ctx iris.Context) {
		// for the sake of simplicity, in order see the logs at the ./_today_.txt
		// ctx.Application().Logger().Infof("Request path: %s", ctx.Path())//显示写入log
		ctx.WriteString("pong")
	})

	// Navigate to http://localhost:8080/ping
	// and open the ./logs{TODAY}.txt file.
	app.Run(
		iris.Addr(":8080"),
		iris.WithoutBanner,
		iris.WithoutServerError(iris.ErrServerClosed),
	)
}
