package main

import (
	"fmt"
	"github.com/spf13/viper"
	"go-web/routes"
	"go-web/utils"
)

func main() {
	r := routes.RegisterRoute()
	utils.InitConfig()
	//model.InitDb()
	if err := r.Run(); err != nil {
		fmt.Printf("startup service failed, err:%v\n", err)
	}
	fmt.Println(viper.Get("mysql.DbHost"))
}
