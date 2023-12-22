package repository

import (
	"GoBookShelf/pkg/models"
	"fmt"

	"github.com/sirupsen/logrus"
)

func GetAllCategoryRepository() ([]models.Category, error) {
	var categories []models.Category
	result := db.Find(&categories)
	if result.Error != nil {
		logrus.WithFields(logrus.Fields{
			"module": "category_repository",
			"errir":  result.Error,
		}).Error("Error retrieving category list")

		return nil, fmt.Errorf("Error retrieving category list:%v ",result.Error)
		

	}
	return categories, nil
}

func CreateCategoryRepository(category models.Category) error {
	result := db.Create(&category)
	if result.Error != nil {
		return result.Error
	}
	return nil

}
