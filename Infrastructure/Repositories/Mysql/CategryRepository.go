package factories

import (
	models "delivery/golang_salesInvoice/Domain/Entities"
	"delivery/golang_salesInvoice/Infrastructure/database"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func (cp *CategoryRepository) New() *CategoryRepository {
	cp.db = database.Connect()
	return cp
}

func (cp CategoryRepository) CreateCategory(category models.Category) *models.Category {
	cp.db.Create(&category)
	return &category
}

func (cp CategoryRepository) FindAll() []models.Category {
	var categories []models.Category
	cp.db.Find(&categories)
	return categories
}

func (cp CategoryRepository) FindById(id int) models.Category {
	var category models.Category
	cp.db.Where("Id = ?", id).Find(&category)
	return category
}

// func (b *Book) CreateBook() *Book {
// 	// db.NewRecord(b)
// 	db.Create(&b)
// 	return b
// }

// func GetAllBooks() []Book {
// 	var books []Book
// 	db.Find(&books)
// 	return books
// }

// func GetBookByID(Id int64) (*Book, *gorm.DB) {
// 	var getBook Book
// 	d := db.Where("ID/?", Id).Find(&getBook)
// 	return &getBook, d
// }

// func DeleteBook(Id int64) Book {
// 	var deletedBook Book
// 	db.Where("ID=?", Id).Delete(&deletedBook)
// 	return deletedBook
// }
