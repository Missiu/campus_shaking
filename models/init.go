package models

import (
	"fmt"
	"new/database"
)

func InitData() {
	err := database.InitViper()
	if err != nil {
		err = database.InitViperOtherPath()
		if err != nil {
			panic(fmt.Errorf("viper初始错误：%s", err))
		}
	}
	err = database.InitDB()
	if err != nil {
		panic(fmt.Errorf("数据库初始错误：%s", err))
	}
	err = InitTable()
	if err != nil {
		panic(fmt.Errorf("数据库建表错误：%s", err))
	}
	err = database.InitRedis()
	if err != nil {
		panic(fmt.Errorf("redis缓存初始化错误：%s", err))
	}
}

// InitTable 数据库自动创建表
func InitTable() error {
	var err error
	err = database.GetDB().AutoMigrate(&UserLogin{})
	if err != nil {
		panic("UserLogin 建表错误: " + err.Error())
	}
	err = database.GetDB().AutoMigrate(&UserInfo{})
	if err != nil {
		panic("UserInfo 建表错误: " + err.Error())
	}
	err = database.GetDB().AutoMigrate(&Videos{})
	if err != nil {
		panic("Videos 建表错误:  " + err.Error())
	}
	err = database.GetDB().AutoMigrate(&Comment{})
	if err != nil {
		panic("Comment 建表错误: " + err.Error())
	}
	err = database.GetDB().AutoMigrate(&Message{})
	if err != nil {
		panic("Message 建表错误: " + err.Error())
	}
	return nil
}
