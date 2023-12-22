package main

import (
	"GoBookShelf/pkg/handler"
	"GoBookShelf/pkg/repository"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	file, err := os.OpenFile("myLogs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatal("Log dosyası açılırken hata oluştu: ", err)
	}

	// Logrus çıktısını dosyaya yönlendir
	logrus.SetOutput(file)
	logrus.SetLevel(logrus.InfoLevel)

	// Örnek kullanım
	logrus.Info("Bu bir bilgi mesajıdır.")

	dataSourceName := "host=localhost port=5432 user=postgres password=12345 dbname=BookDatabase sslmode=disable"
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed Database connection :%v", err)
	}
	repository.InitializeDB(db)
	// repository.InitializeDB(dataSourceName)
	router := gin.Default()

	//router tanımlamaları
	handler.InitializeRoutes(router)

	//Run server

	err = router.Run(":8080")
	if err != nil {
		log.Fatalf("Sunucu başlatılırken hata: %v", err)
	}

}
