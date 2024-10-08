package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/stupidlamo/songlist/backend/internal/auth"
	"github.com/stupidlamo/songlist/backend/internal/controllers"
	"github.com/stupidlamo/songlist/backend/internal/database"
)

func main() {
	r := gin.Default()

	// Подключение к базе данных
	database.Connect()

	// Маршруты
	public := r.Group("/api")
	{
		public.POST("/login", controllers.Login)
		public.POST("/register", controllers.Register)
	}

	protected := r.Group("/api")
	protected.Use(auth.JWTAuthMiddleware())
	{
		protected.GET("/songlists", controllers.GetSongLists)
		protected.POST("/songlists", controllers.CreateSongList)
		protected.PUT("/songlists/:id", controllers.UpdateSongList)
		protected.DELETE("/songlists/:id", controllers.DeleteSongList)
	}

	// Запуск сервера
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
