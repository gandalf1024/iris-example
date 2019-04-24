package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/kataras/golog"
)

var DB *gorm.DB

//初始化连接
func init() {
	DB, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/wns")
	if err != nil {
		panic("连接数据库失败")
	}
	golog.Info("数据库连接成功")
}

func Close() {
	if DB != nil {
		err := DB.Close()
		if err != nil {
			golog.Error("数据库关闭连接失败")
		}
	}
}
