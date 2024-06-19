package main

import (
	"fmt"
	"new/models"
	"new/routes"
)

func main() {
	models.InitData()
	r := routes.Router()
	err := r.Run()
	if err != nil {
		panic(fmt.Errorf("启动失败: %s", err))
	}
}
