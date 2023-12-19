package main

import (
	"fmt"
	"go-web/routes"
)

func main() {
	r := routes.RegisterRoute()
	if err := r.Run(); err != nil {
		fmt.Printf("startup service failed, err:%v\n", err)
	}
}
