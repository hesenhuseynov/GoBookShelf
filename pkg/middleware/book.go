package middleware

// func BooksMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		logrus.WithFields(logrus.Fields{
// 			"method": c.Request.Method,
// 			"path":   c.Request.URL.Path,
// 			"time":   time.Now(),
// 			"os":     c.Request.UserAgent(),
// 		}).Info("Reseceived  book request on middleware")

// 		// Burada istek üzerinde bazı kontroller yapılabilir
// 		// Örneğin, bazı header'lar veya parametreler kontrol edilebilir

// 		c.Next() // Sonraki handler'a geçiş
// 	}
// }
