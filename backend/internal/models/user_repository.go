package repositories

import (
	"github.com/yourusername/yourproject/internal/models"
	"gorm.io/gorm"
)

func FindUserByUsername(db *gorm.DB, username string) (*models.User, error) {
	var user models.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateUser(db *gorm.DB, user *models.User) error {
	return db.Create(user).Error
}
