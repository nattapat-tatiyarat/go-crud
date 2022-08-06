package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func connect() {
	dsn := "user:password@tcp(127.0.0.1:3306)/intern?charset=utf8mb4&parseTime=True&loc=Local"
	mySql, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db = mySql
	if err != nil {
		panic("failed to connect database")
	}
}
