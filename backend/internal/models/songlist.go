package models

import "gorm.io/gorm"

type SongList struct {
	gorm.Model
	Name     string `json:"name"`
	UserName string `json:"username"`
	Songs    []Song `json:"songs"`
}

type Song struct {
	gorm.Model
	Title      string `json:"title"`
	Artist     string `json:"artist"`
	SongListID uint   `json:"songListId"`
}
