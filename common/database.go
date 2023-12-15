package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go-superset/model"
)

func GetDB() *gorm.DB {
	driverName := "mysql"
	host := "localhost"
	port := "3306"
	database := "godemo"
	username := "root"
	password := "abcd1234!"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("fail to connect database, err: " + err.Error())
	}

	// 自动创建数据表
	db.AutoMigrate(&model.User{})
	return db
}
