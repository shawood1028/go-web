package model

import (
	"fmt"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

// InitMysqlDb 初始化mysql
func InitMysqlDb() {

	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local")
}
