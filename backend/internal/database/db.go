package database

import (
	"log"

	"github.com/yourusername/yourproject/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=youruser password=yourpassword dbname=yourdbname port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Миграция моделей
	db.AutoMigrate(&models.User{}, &models.SongList{}, &models.Song{})

	DB = db
}
