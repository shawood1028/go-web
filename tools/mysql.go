package tools

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var _db *gorm.DB

func Init() {
	DbUser := viper.Get("mysql.DbUser")
	DbPassWord := viper.Get("mysql.DbPassWord")
	DbHost := viper.GetString("mysql.DbHost")
	DbPort := viper.Get("mysql.DbPort")
	DbName := viper.Get("mysql.DbName")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DbUser,
		DbPassWord,
		DbHost,
		DbPort,
		DbName,
	)
	var err error
	_db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	sqlDB, _ := _db.DB()

	// TODO 设置数据库连接池参数
	sqlDB.SetMaxOpenConns(100) // 设置数据库连接池最大连接数
	fmt.Println("MySQL初始化")
}

func GetMysqlDB() *gorm.DB {
	return _db
}
