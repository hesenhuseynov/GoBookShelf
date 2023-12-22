package handler

import (
	"GoBookShelf/pkg/models"
	"GoBookShelf/pkg/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

func InitializeRoutes(router *gin.Engine) {
	router.GET("/books", getBooksHandler)
	router.POST("/books", createBookHandler)
	router.DELETE("/books/:id", deleteBookHandler)
	router.PUT("/books/:id", updateBookHandler)
	router.GET("/books/:id", getBookByIdHandler)
	router.GET("/categories", getCategoriesHandler)
	router.POST("/categories", createCategoryHandler)
	router.GET("/books/category/:category_id", getBooksByCategoryHandler)

	//GET /authors/{id}/books: Belirli bir yazarın kitaplarını listeler.

}

// getBooks : Tüm kitapları döndürür
func getBooksHandler(c *gin.Context) {

	// Kitapları getiren servis kodları...
	books, err := service.GetAllBooksService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)

}

func updateBookHandler(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")
	err := service.UpdateBookService(id, book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book updated", "book": book})

}

func deleteBookHandler(c *gin.Context) {
	id := c.Param("id")
	err := service.DeleteBookService(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errir": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}

func createBookHandler(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := service.CreateBookService(book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	c.JSON(http.StatusCreated, gin.H{"message": "Book Added", "book": book})

}

func getBooksByCategoryHandler(c *gin.Context) {
	// Girdi Doğrulaması: URL parametresini integer'a dönüştürme
	CategoryIDStr := c.Param("category_id")
	categoryID, err := strconv.Atoi(CategoryIDStr)

	// categoryID, err := strconv.Atoi(c.Param(CategoryIDStr))
	if err != nil {
		// Hata Yönetimi: Geçersiz parametre durumunda hata mesajı gönder
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}
	// Veritabanı sorgusu: SQL Enjeksiyonu Önleme
	books, err := service.GetBooksByCategoryIDService(categoryID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"module": "book_handler",
			"error":  err.Error(),
			"action": "Retrieving books by category",
		}).Error("Error in getBooksByCategoryHandler")
		// Hata Yönetimi: İç sunucu hatası durumunda hata mesajı gönder
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return

	}
	c.JSON(http.StatusOK, books)

}
func getBookByIdHandler(c *gin.Context) {
	bookIDStr := c.Param("id")
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bookId"})
		return
	}
	book, err := service.GetBookByIdService(bookID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"module": "book_handler",
			"error":  err.Error(),
			"action": "Retrieving book by ID",
		}).Error("Error in getBookByIDHandler")

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return

	}
	c.JSON(http.StatusOK, book)

}
