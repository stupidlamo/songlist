package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/stupidlamo/songlist/backend/internal/models"
	"github.com/stupidlamo/songlist/backend/internal/services"
)

func GetSongLists(c *gin.Context) {
	username, _ := c.Get("username")
	songLists, err := services.GetSongLists(username.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, songLists)
}

func CreateSongList(c *gin.Context) {
	var songList models.SongList
	if err := c.ShouldBindJSON(&songList); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	username, _ := c.Get("username")
	songList.UserName = username.(string)
	if err := services.CreateSongList(&songList); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, songList)
}

func UpdateSongList(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var songList models.SongList
	if err := c.ShouldBindJSON(&songList); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.UpdateSongList(id, &songList); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, songList)
}

func DeleteSongList(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := services.DeleteSongList(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Song list deleted successfully"})
}
