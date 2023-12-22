package repository

import (
	"GoBookShelf/pkg/models"
	"errors"
	"fmt"
	"log"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitializeDB(database *gorm.DB) {
	db = database
	err := db.AutoMigrate(&models.Book{})
	if err != nil {
		log.Fatalf("Automigrated failed:%v", err)

	}

}
func GetAllBooksRepository() ([]models.Book, error) {
	var books []models.Book
	result := db.Find(&books)
	if result.Error != nil {

		return nil, result.Error
	}
	return books, nil
}

func GetBookByIdRepository(bookID int) (models.Book, error) {
	var book models.Book
	result := db.First(&book, bookID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return models.Book{}, fmt.Errorf("Book not Founded:%d", bookID)
		}
		return models.Book{}, result.Error

	}
	return book, nil

}

func GetBooksByCategoryID(categoryID int) ([]models.Book, error) {
	// Fonksiyonun içeriği
	var books []models.Book
	result := db.Where("category_id=?", categoryID).Find(&books)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("Books not found in category: %d", categoryID)

		}
		return nil, result.Error
	}

	if len(books) == 0 {
		// Kayıt bulunamadı durumu için özel hata loglama
		logrus.WithFields(logrus.Fields{
			"module": "book_repository",
			"error":  "No books found",
			"action": "Retrieving books by category",
		}).Warn("No books found for category ID:", categoryID)
		return nil, fmt.Errorf("No books found for category ID: %d", categoryID)
	}
	//result := db.Where(book.CategoryID == &categoryID).First(&book)

	return books, nil

}

func CreateBookRepository(book models.Book) error {
	result := db.Create(&book)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateBookRepository(id string, book models.Book) error {
	var dbBook models.Book
	result := db.First(&dbBook, id)
	if result.Error != nil {
		return result.Error
	}
	dbBook.Title = book.Title
	// dbBook.AuthorID = book.AuthorID
	result = db.Save(&dbBook)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteBookRepository(id string) error {
	result := db.Delete(&models.Book{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
