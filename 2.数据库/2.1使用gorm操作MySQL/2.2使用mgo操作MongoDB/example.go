package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

type User struct {
	// ID bson.ObjectId `bson:"_id"`
	Name      string    `bson:"name"`
	Age       int       `bson:"age"`
	JonedAt   time.Time `bson:"joned_at"`
	Interests []string  `bson:"interests"`
}

func main() {
	//连接mongo
	session, err := mgo.Dial("")
	if err != nil {
		fmt.Println("数据库连接失败")
		panic(err)
	} else {
		fmt.Println("数据库连接成功")
	}
	defer session.Close()
	// 连接数据库
	session.SetMode(mgo.Monotonic, true)
	db := session.DB("iris")

	//指向某一个集合（表）
	usersCollection := db.C("users")
	//插入数据
	err = usersCollection.Insert(&User{
		Name:      "gbs",
		Age:       12,
		JonedAt:   time.Now(),
		Interests: []string{"eating", "sleeping"},
	})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("数据插入成功")
	}
	//查找数据
	var users []User
	usersCollection.Find(nil).All(&users) //将数据放进 users中
	fmt.Println(users)

	usersCollection.Find(bson.M{"name": "gbs"}).All(&users) //有条件
	fmt.Println(users)

	usersCollection.Update(bson.M{"name": "gbs"},
		bson.M{"$set": bson.M{
			"name": "great", "age": 13,
		}})
}
