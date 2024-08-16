package repositories

import (
	"github.com/yourusername/yourproject/internal/models"
	"gorm.io/gorm"
)

func FindSongListsByUsername(db *gorm.DB, username string) ([]models.SongList, error) {
	var songLists []models.SongList
	if err := db.Where("user_name = ?", username).Find(&songLists).Error; err != nil {
		return nil, err
	}
	return songLists, nil
}

func CreateSongList(db *gorm.DB, songList *models.SongList) error {
	return db.Create(songList).Error
}

func UpdateSongList(db *gorm.DB, id int, songList *models.SongList) error {
	return db.Model(&models.SongList{}).Where("id = ?", id).Updates(songList).Error
}

func DeleteSongList(db *gorm.DB, id int) error {
	return db.Delete(&models.SongList{}, id).Error
}
