package routes

import (
	"GoBookShelf/pkg/handler"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine, user *handler.UserHandler) {

	router.GET("/books", handler.GetBooksHandler)
	router.POST("/books", handler.CreateBookHandler)
	router.DELETE("/books/:id", handler.DeleteBookHandler)
	router.PUT("/books/:id", handler.UpdateBookHandler)
	router.GET("/books/:id", handler.GetBookByIdHandler)
	router.GET("/books/language", handler.GetBooksByLanguageHandler)
	router.GET("/categories", handler.GetCategoriesHandler)
	router.POST("/categories", handler.CreateCategoryHandler)
	router.GET("/books/category/:category_id", handler.GetBooksByCategoryHandler)
	// Category routes gruplaması
	// User routes (UserHandler yapısı üzerinden)
	userRoutes := router.Group("/users")
	{   
		userRoutes.POST("/", user.Register)
		userRoutes.POST("/login",user.Login)
		// ... diğer user route'ları eklenebilir ...
	}
}
