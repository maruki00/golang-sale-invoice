package db

import (
	"gorm.io/gorm"
)

var _db *gorm.DB

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	// db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookByID(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	d := db.Where("ID/?", Id).Find(&getBook)
	return &getBook, d
}

func DeleteBook(Id int64) Book {
	var deletedBook Book
	db.Where("ID=?", Id).Delete(&deletedBook)
	return deletedBook
}
