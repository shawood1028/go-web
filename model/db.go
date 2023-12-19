package model

import (
	"fmt"
	"go-web/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"os"
)

var db *gorm.DB
var err error

func InitDb() {
	var dbconfig = utils.InitConfig().MySQL
	fmt.Println(dbconfig.DbHost)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbconfig.DbUser,
		dbconfig.DbPassWord,
		dbconfig.DbHost,
		dbconfig.DbPort,
		dbconfig.DbName,
	)
	fmt.Println(dsn)
	// TODO 初始化数据库
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// gorm日志模式：silent
		Logger: logger.Default.LogMode(logger.Silent),
		// 外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
		// 禁用默认事务（提高运行速度）
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Println("连接数据库失败，请检查参数：", err)
		os.Exit(1)
	}
}
