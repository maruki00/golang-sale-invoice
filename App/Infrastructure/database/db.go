package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = nil
)

type DBHandler struct {
	db *gorm.DB
}

func Connect() *gorm.DB {

	if db == nil {
		d, err := gorm.Open(mysql.Open("sale_invoce_api:user@tcp(127.0.0.1:3306)/sale_invoce_api?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
		if err != nil {
			// log.Fatal("DB Error : ", err.Error())
			panic(err)
		}
		db = d
	}
	return db
}

func NewDB() *DBHandler {
	return &DBHandler{
		db: Connect(),
	}
}

// func (dbhandler *DBHandler) Create(object interface{}) {
// 	dbhandler.db.Create(object)
// }

// func (dbhandler *DBHandler) Update(object interface{}) {
// 	dbhandler.db.Save(object)
// }

// func (dbhandler *DBHandler) Delete(object interface{}) {
// 	dbhandler.db.Delete(object)
// }

// func (dbhandler *DBHandler) Find(object interface{}){
// 	dbhandler.db.Find(object)
// }

// func (dbhandler *DBHandler) Where(object interface{}) *gorm.DB{
// 	dbhandler.db.
// 	return dbhandler.db.Where(object)
// }
