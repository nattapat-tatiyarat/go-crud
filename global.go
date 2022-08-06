package main

import (
	"gorm.io/gorm"
)

var db *gorm.DB
var todos []Todos

type Todos struct {
	gorm.Model
	Task string
}

type RequestBody struct {
	Id   int    `json:"id"`
	Task string `json:"task"`
}
