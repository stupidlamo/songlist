package services

import (
	"github.com/stupidlamo/songlist/backend/internal/database"
	"github.com/stupidlamo/songlist/backend/internal/models"
	"github.com/stupidlamo/songlist/backend/internal/repositories"
)

// Получение всех списков песен для пользователя
func GetSongLists(username string) ([]models.SongList, error) {
	return repositories.FindSongListsByUsername(database.DB, username)
}

// Создание нового списка песен
func CreateSongList(songList *models.SongList) error {
	return repositories.CreateSongList(database.DB, songList)
}

// Обновление списка песен
func UpdateSongList(id int, songList *models.SongList) error {
	return repositories.UpdateSongList(database.DB, id, songList)
}

// Удаление списка песен
func DeleteSongList(id int) error {
	return repositories.DeleteSongList(database.DB, id)
}
