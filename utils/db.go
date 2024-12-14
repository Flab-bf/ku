package utils

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	dns := "root:dzf244106_F@tcp(127.0.0.1:3306)/mytest?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败: " + err.Error())
	}
}
