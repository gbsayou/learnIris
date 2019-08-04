# 使用 gorm 操作 MySQL

---

gorm 是 Golang 写的，开发人员友好的 ORM 库。
更加详细的介绍请看[文档](http://jinzhu.me/gorm/)

## 1. 定义模型

使用 Go 的结构体定义 gorm 的模型。以下代码简单示范了三个模型，并定义了各个模型之间的关系。
gorm.Model 表示使用 gorm 提供的基本字段，包括 ID，CreatedAt，UpdatedAt，DeletedAt。

```Go
type User struct {
    gorm.Model
    Birthday     time.Time
    Age          int
    Name         string  `gorm:"size:255"`       // string 的默认长度为255, 使用 tag 可自定义。
    Num          int     `gorm:"AUTO_INCREMENT"` // 自增长
}

type Email struct {
    ID      int
    UserID  int     `gorm:"index"` // 外键, tag `index`是为该列创建索引
    Email   string  `gorm:"type:varchar(100);unique_index"` // `type`设置字段类型, `unique_index` 表示为该列设置唯一索引
}

type CreditCard struct {
    gorm.Model
    UserID  uint
    Number  string
}

```

## 2. 基本 CRUD

```Go
var user = User{Birthday: time.Now(),Age:1,Name:"gbs"}
db.Create(&user)
// insert into users (Birthday,Age,Name) values()

var users []User
db.Where(&User{Name: "gbs", Age: 20}).Find(&user)
// select * from users where name = 'gbs' and age = 20

db.Model(&user).Where("id = ", 1).Update("name", "gbs")
// update users set name = 'gbs' where id = 1

db.Where("id = ?", 1).Delete(User{})
// delete from users where id = 1
```

针对删除，如果表中定义了 DeletedAt 字段，则删除的时候，是执行软删除。

## 3.使用事务

先利用 gorm 开始一个事务，然后在这个事务中执行具体的操作。当发生错误时，回滚事务；当操作进行顺利，则需要提交事务。

```Go
func UserRegister(ctx iris.Context){
    trx := db.Begin()//开始并获取一个事务
    name := ctx.Params().Get("name")//从请求中获取name
    //使用当前事务执行数据库操作
    if err:= trx.Create(&User{Name:name}).Error;err!=nil{
        trx.Rollback()
        panic("用户注册失败")
    }
    trx.Commit()
    ctx.JSON("用户注册成功")
}
```
