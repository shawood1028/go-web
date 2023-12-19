package main

import (
	"fmt"
	"go-web/config"
	"go-web/routes"
)

func main() {
	config.InitConfig()
	r := routes.RegisterRoute()
	if err := r.Run(); err != nil {
		fmt.Printf("startup service failed, err:%v\n", err)
		fmt.Printf("3333")
		fmt.Printf("44444")
	}

}
