package services

import (
	"errors"

	"github.com/yourusername/yourproject/internal/auth"
	"github.com/yourusername/yourproject/internal/database"
	"github.com/yourusername/yourproject/internal/models"
	"github.com/yourusername/yourproject/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

// Создание нового пользователя
func CreateUser(user *models.User) error {
	// Хеширование пароля перед сохранением
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return repositories.CreateUser(database.DB, user)
}

// Логин пользователя
func LoginUser(user *models.User) (string, error) {
	// Поиск пользователя по username
	foundUser, err := repositories.FindUserByUsername(database.DB, user.Username)
	if err != nil {
		return "", errors.New("user not found")
	}

	// Проверка пароля
	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password))
	if err != nil {
		return "", errors.New("invalid password")
	}

	// Генерация JWT токена
	token, err := auth.GenerateJWT(foundUser.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}
