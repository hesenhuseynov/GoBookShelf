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


// getBooksHandler - Tüm Kitapları Getir
// @Summary Tüm kitapları listeler
// @Description Tüm kitapları döndürür
// @Tags books
// @Accept json
// @Produce json
// @Success 200 {array} models.Book
// @Router /books [get]
func GetBooksHandler(c *gin.Context) {

	// Kitapları getiren servis kodları...
	books, err := service.GetAllBooksService()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
			"where": "getBooksHandler",
		}).Error("Failed to get all books")

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// logrus.Info("Retrieved all books successfully.")
	c.JSON(http.StatusOK, books)

}

// updateBookHandler - Kitap Güncelle
// @Summary Belirtilen ID'ye sahip kitabı günceller
// @Description ID ile kitap bilgilerini günceller
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Kitap ID"
// @Param book body models.Book required "Kitap Bilgileri"
// @Success 200 {object} Book
// @Router /books/{id} [put]
func UpdateBookHandler(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
			"where": "updateBookHandler",
		}).Error("Bad request for book update")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")
	err := service.UpdateBookService(id, book)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"bookID": id,
			"error":  err.Error(),
			"where":  "updateBookHandler",
		}).Error("Failed to update book")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	logrus.WithFields(logrus.Fields{
		"bookID": id,
	}).Info("Updated book successfully.")
	c.JSON(http.StatusOK, gin.H{"message": "Book updated", "book": book})

}

// deleteBookHandler - Kitap Sil
// @Summary Belirtilen ID'ye sahip kitabı siler
// @Description ID ile kitap siler
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Kitap ID"
// @Success 200 {string} string "Book deleted"
// @Router /books/{id} [delete]
func DeleteBookHandler(c *gin.Context) {
	id := c.Param("id")
	err := service.DeleteBookService(id)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"bookID": id,
			"error":  err.Error(),
			"where":  "deleteBookHandler",
		}).Error("Failed to delete book")
		c.JSON(http.StatusInternalServerError, gin.H{"errir": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}

// createBookHandler - Kitap Ekle
// @Summary Yeni bir kitap ekler
// @Description Yeni kitap oluşturur
// @Tags books
// @Accept json
// @Produce json
// @Param book body models.Book required "Kitap Bilgileri"
// @Success 201 {object} Book
// @Router /books [post]

func CreateBookHandler(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
			"where": "createBookHandler",
		}).Error("Bad request for book creation")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(book); err != nil {

		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
			"where": "createBookHandler",
		}).Error("Validation failed for book creation")

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := service.CreateBookService(book)
	if err != nil {

		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
			"where": "createBookHandler",
		}).Error("Failed to create book")

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	c.JSON(http.StatusCreated, gin.H{"message": "Book Added", "book": book})

}

// getBooksByCategoryHandler - Kategoriye Göre Kitapları Listele
// @Summary Belirli bir kategorideki kitapları listeler
// @Description Kategori ID'sine göre kitapları getirir
// @Tags books
// @Accept json
// @Produce json
// @Param category_id path int true "Kategori ID"
// @Success 200 {array} models.Book
// @Router /books/category/{category_id} [get]
func GetBooksByCategoryHandler(c *gin.Context) {
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
			"where":  "getBooksByCategoryHandler",
			"action": "Retrieving books by category",
		}).Error("Error in getBooksByCategoryHandler")
		// Hata Yönetimi: İç sunucu hatası durumunda hata mesajı gönder
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return

	}
	c.JSON(http.StatusOK, books)

}

// getBookByIdHandler - Kitabı ID'ye Göre Getir
// @Summary Belirtilen ID'ye sahip kitabı getirir
// @Description ID ile kitap bilgilerini getirir
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Kitap ID"
// @Success 200 {object} models.Book
// @Router /books/{id} [get]
func GetBookByIdHandler(c *gin.Context) {
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
			"where":  "getBookByIdHandler",
			"action": "Retrieving book by ID",
		}).Error("Error in getBookByIDHandler")

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return

	}
	c.JSON(http.StatusOK, book)

}

// getBooksByLanguageHandler - Dil Bazında Kitapları Listele
// @Summary Belirli bir dildeki kitapları listeler
// @Description Sorgu parametresi olarak verilen dildeki kitapları getirir
// @Tags books
// @Accept json
// @Produce json
// @Param language query string true "Dil"
// @Success 200 {array} models.Book
// @Router /books/language [get]
func GetBooksByLanguageHandler(c *gin.Context) {
	language := c.Query("language")                           // Dil bilgisini sorgu parametresinden al
	books, err := service.GetBooksByLanguageService(language) // Servis katmanını çağır
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"module":   "book_handler",
			"error":    err.Error(),
			"where":    "getBooksByLanguageHandler",
			"language": language,
			"action":   "Retrieving books by language",
		}).Error("Error in getBooksByLanguageHandler")

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving books by language: " + language})
		return
	}

	if len(books) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No books found for the specified language: " + language})
		return
	}

	c.JSON(http.StatusOK, books)
}
