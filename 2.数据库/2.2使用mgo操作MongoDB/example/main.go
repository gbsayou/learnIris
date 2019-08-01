package main

import (
	"./models"
	"fmt"
	"github.com/kataras/iris"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"time"
)

func main() {
	app := iris.Default()
	//连接mongo
	session, err := mgo.Dial("")
	if err != nil {
		fmt.Println("数据库连接失败")
		panic(err)
	} else {
		fmt.Println("数据库连接成功")
	}
	defer session.Close()
	// 指向数据库
	session.SetMode(mgo.Monotonic, true)
	db := session.DB("iris")
	app.Post("/user", func(ctx iris.Context) {
		usersCollection := db.C("users")
		err = usersCollection.Insert(&models.User{
			Name:      ctx.FormValueDefault("name", "great"),
			Age:       12,
			Interests: make([]string, 0),
			JonedAt:   time.Now(),
		})
		if err != nil {
			ctx.JSON(iris.Map{
				"status": "插入数据失败",
				"error":  err,
			})
		} else {
			ctx.Writef("插入成功")
		}
	})
	app.Get("/user/{name}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
		var results []models.User
		usersCollection := db.C("users")
		usersCollection.Find(bson.M{"name": name}).All(&results)
		ctx.JSON(results)
	})
	app.Put("/user/{name}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
		age, _ := strconv.Atoi(ctx.FormValueDefault("age", "15"))
		usersCollection := db.C("users")
		result := usersCollection.Update(bson.M{"name": name},
			bson.M{"$set": bson.M{
				"age": age,
			}})
		ctx.JSON(result)
	})
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
