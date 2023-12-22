package service

import (
	"GoBookShelf/pkg/models"
	"GoBookShelf/pkg/repository"
)

func GetAllCategoryService() ([]models.Category, error) {
	return repository.GetAllCategoryRepository()

}
func CreateCategoryService(category models.Category) error {
	return repository.CreateCategoryRepository(category)
}

