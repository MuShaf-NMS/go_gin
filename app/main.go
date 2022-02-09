package app

import (
	"fmt"
	"log"
	"os"

	"github.com/MuShaf-NMS/go_gin/app/config"
	"github.com/MuShaf-NMS/go_gin/app/models"
	"github.com/MuShaf-NMS/go_gin/app/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func getEnv(key string, defaultVal string) string {
	if value, ok := os.LookupEnv(key); ok {
		fmt.Println(value)
		return value
	}
	return defaultVal
}

func InitializeDB(config config.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DB_User, config.DB_Pass, config.DB_Host, config.DB_Port, config.DB_Name)
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("DB connection failed")
	} else {
		println("DB connction success")
	}
	errMigrate := connection.AutoMigrate(models.User{}, models.Todo{})
	if errMigrate != nil {
		panic(errMigrate)
	}
	return connection
}

func CloseDB(connction *gorm.DB) {
	db, err := connction.DB()
	if err != nil {
		panic("Failed to close DB connection")
	}
	db.Close()
}

func Run() {
	config := config.Config{
		App_Port:  getEnv("APP_PORT", "8000"),
		Env:       getEnv("APP_ENV", "development"),
		SecretKey: getEnv("SECRET", "go_gin_test"),
		DB_Host:   getEnv("DB_HOST", "localhost"),
		DB_User:   getEnv("DB_USER", "root"),
		DB_Pass:   getEnv("DB_PASS", "root"),
		DB_Name:   getEnv("DB_NAME", "gin_test"),
		DB_Port:   getEnv("DB_PORT", "3306"),
	}

	r := gin.Default()
	db := InitializeDB(config)
	// defer CloseDB(db)
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error .env file")
	}

	r.RedirectTrailingSlash = true
	router.RoutingAuth(r, db, config)
	router.RoutingUser(r, db, config)
	router.RoutingTodo(r, db, config)

	fmt.Println(config)

	r.Run(":" + config.App_Port)
}
