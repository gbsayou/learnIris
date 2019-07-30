package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Birthday time.Time
	Age      int
	Name     string `gorm:"size:255"`       // string 的默认长度为255, 使用 tag 可自定义。
	Num      int    `gorm:"AUTO_INCREMENT"` // 自增长
}
type Email struct {
	ID     int
	UserID int    `gorm:"index"`                          // 外键, tag `index`是为该列创建索引
	Email  string `gorm:"type:varchar(100);unique_index"` // `type`设置字段类型, `unique_index` 表示为该列设置唯一索引
}

type CreditCard struct {
	gorm.Model
	UserID uint
	Number string
}
