package tools

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var _db *gorm.DB

func init() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		fmt.Println("Config file changed:", e.Name)
	})
	workDir, _ := os.Getwd()
	viper.SetConfigName("dev")
	viper.SetConfigType("toml")
	viper.AddConfigPath(workDir + "/etc")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("配置文件初始化")
	return
}

func GetMysqlDB() *gorm.DB {
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
	} else {
		sqlDB, _ := _db.DB()
		// TODO 设置数据库连接池参数
		sqlDB.SetMaxOpenConns(100) // 设置数据库连接池最大连接数
		fmt.Println("MySQL初始化成功")
	}
	return _db
}
