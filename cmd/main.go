package main

import (
	"GoBookShelf/pkg/handler"
	"GoBookShelf/pkg/models"
	"GoBookShelf/pkg/repository"
	"GoBookShelf/pkg/service"
	"GoBookShelf/routes"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nullseed/logruseq"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	file, err := os.OpenFile("myLogs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatal("Log dosyası açılırken hata oluştu: ", err)
	}

	// Logrus çıktısını dosyaya ve Seq'e yönlendir
	logrus.SetOutput(file)
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// Seq için yeni bir hook oluşturun (Seq sunucu URL ve API anahtarınızı buraya girin)
	seqHook := logruseq.NewSeqHook("http://localhost:5341", logruseq.OptionAPIKey("TGKfiXrjKWzoZr5MBG42"))

	if err != nil {
		logrus.Fatal("Seq hook oluşturulurken hata: ", err)
	}

	// Logrus'a Seq hook'unu ekleyin
	logrus.AddHook(seqHook)
	logrus.SetLevel(logrus.InfoLevel)

	// Örnek kullanım+
	// logrus.Info("Bu bir bilgi mesajıdır.")

	// Veritabanı bağlantısı ve router ayarları
	dataSourceName := "host=localhost port=5432 user=postgres password=12345 dbname=BookDatabase sslmode=disable"
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		logrus.Fatalf("Veritabanı bağlantısı başarısız: %v", err)
	}
	repository.InitializeDB(db)

	router := gin.Default()
	// jwtService := service.NewJWTService()
	// middleware.Authenticate(jwtService)
	router.Use()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:4200"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type"}
	router.Use(cors.New(config))
	// handler.InitializeRoutes(router)
	// UserRepository örneğini oluşturma
	userRepo := repository.NewUserRepository(db)

	// UserService örneğini oluşturma
	userService := service.NewUserService(userRepo)

	// UserHandler örneğini oluşturma
	userHandler := handler.NewUserHandler(userService)
	routes.InitializeRoutes(router, userHandler)

	err = router.Run(":8080")
	if err != nil {
		logrus.Fatalf("Sunucu başlatılırken hata: %v", err)
	}
	
	if err := db.AutoMigrate(&models.Book{}, models.User{}); err != nil {
		logrus.Fatalf("Automigrate error :%v", err)
	}
}
