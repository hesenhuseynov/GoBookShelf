package service

import (
	"GoBookShelf/pkg/models"
	"GoBookShelf/pkg/repository"
)

func GetAllBooksService() ([]models.Book, error) {
	return repository.GetAllBooksRepository()
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
