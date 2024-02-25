package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = nil
)

func Connect() {
	if db == nil {
		d, err := gorm.Open(mysql.Open("sale_invoce_api:user@tcp(127.0.0.1:3306)/sale_invoce_api?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
		if err != nil {
			// log.Fatal("DB Error : ", err.Error())
			panic(err)
		}
		db = d
	}
}

func GetDB() *gorm.DB {
	return db
}
