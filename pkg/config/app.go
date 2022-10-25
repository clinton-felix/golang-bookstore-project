package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// the idea of the app.go file is to create a var called DB which will help other
// parts of the application to interact with the database...

var (
	db * gorm.DB
)

// function helps us open a connection with the database (mysql DB in this instance)
func Connect()  {
	d, err := gorm.Open("mysql", "root:Your_Password@tcp(127.0.0.1:3306)/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}