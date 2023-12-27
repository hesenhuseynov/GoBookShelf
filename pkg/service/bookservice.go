package service

import (
	"GoBookShelf/pkg/models"
	"GoBookShelf/pkg/repository"

	"github.com/sirupsen/logrus"
)

func GetAllBooksService() ([]models.Book, error) {
	books, err := repository.GetAllBooksRepository()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service": "GetAllBooksService",
			"error":   err.Error(),
		}).Error("Error getting all books")
	}
	return books, err

}

func UpdateBookService(id string, book models.Book) error {

	return repository.UpdateBookRepository(id, book)
}

func CreateBookService(book models.Book) error {
	return repository.CreateBookRepository(book)
}
func DeleteBookService(id string) error {
	return repository.DeleteBookRepository(id)
}
func GetBooksByCategoryIDService(categoryId int) ([]models.Book, error) {
	return repository.GetBooksByCategoryID(categoryId)
}
func GetBookByIdService(bookId int) (models.Book, error) {
	return repository.GetBookByIdRepository(bookId)
}
func GetBooksByLanguageService(language string) ([]models.Book, error) {
	return repository.GetBooksByLanguageRepository(language)
}
