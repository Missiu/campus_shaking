package database

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("mysql.username"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.database"),
	)
	//fmt.Println("dsn:", dsn)
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("数据库链接错误: %w", err)
	}
	db = DB
	return nil
}

func GetDB() *gorm.DB {
	if db == nil {
		err := InitViperOtherPath()
		if err != nil {
			err = InitViper()
			if err != nil {
				panic(fmt.Errorf("初始化Viper失败: %s", err))
			}
		}
		err = InitDB()
		if err != nil {
			panic(fmt.Errorf("初始化mysql失败: %s", err))
		}
	}
	return db
}
