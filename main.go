package main

import (
	"final-task-pbi-rakamin-fullstack-m.aldi_gunawan/handlers"
	"final-task-pbi-rakamin-fullstack-m.aldi_gunawan/models"
	"github.com/gin-gonic/gin"
)

func main() {
	// Koneksi ke database
	models.ConnectDatabase()

	// Setup Gin router
	router := gin.Default()

	// Middleware untuk autentikasi JWT
	router.Use(handlers.AuthMiddleware())

	// Route untuk endpoint pengguna
	router.POST("/users/register", handlers.RegisterUserHandler)
	router.POST("/users/login", handlers.LoginUserHandler)
	router.PUT("/users/:userId", handlers.UpdateUserHandler)
	router.DELETE("/users/:userId", handlers.DeleteUserHandler)

	// Route untuk endpoint foto
	router.POST("/photos", handlers.CreatePhotoHandler)
	router.GET("/photos", handlers.GetPhotosHandler)
	router.PUT("/photos/:photoId", handlers.UpdatePhotoHandler)
	router.DELETE("/photos/:photoId", handlers.DeletePhotoHandler)

	// Menjalankan server
	router.Run(":8080")
}
