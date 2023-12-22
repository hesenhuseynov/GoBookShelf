package handler

import (
	"GoBookShelf/pkg/models"
	"GoBookShelf/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

func getCategoriesHandler(c *gin.Context) {
	categories, err := service.GetAllCategoryService()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"module": "category_handler",
			"error":  err.Error(),
			"action": "Retrieving category list",
		}).Error("Error in getCategoriesHandler")

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return

	}
	c.JSON(http.StatusOK, categories)

}

func createCategoryHandler(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := service.CreateCategoryService(category)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"module": "category_handler",
			"error":  err.Error(),
			"action": "Creating Category ",
		}).Error("error in CreateCategoryService")

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Category Added", "category": category})

}
