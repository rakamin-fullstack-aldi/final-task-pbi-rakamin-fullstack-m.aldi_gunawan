package handlers

import (
	"net/http"
	"net/url"

	"final-task-pbi-rakamin-fullstack-m.aldi_gunawan/models"
	"github.com/gin-gonic/gin"
)

// CreatePhotoHandler membuat photo baru
func CreatePhotoHandler(c *gin.Context) {
	var newPhoto models.Photo
	if err := c.ShouldBindJSON(&newPhoto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Pastikan URL foto tidak kosong
	if newPhoto.PhotoUrl == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "PhotoUrl is required"})
		return
	}

	// Validasi URL foto
	if _, err := url.ParseRequestURI(newPhoto.PhotoUrl); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid PhotoUrl format"})
		return
	}

	// Pastikan user yang sedang login
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userIDUint, _ := userID.(uint)
	newPhoto.UserID = userIDUint

	// Simpan photo ke database
	if err := models.DB.Create(&newPhoto).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create photo"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Photo created successfully", "photo": newPhoto})
}

// UpdatePhotoHandler mengupdate photo berdasarkan photo ID
func UpdatePhotoHandler(c *gin.Context) {
	photoID := c.Param("photoId")

	var photo models.Photo
	if err := models.DB.First(&photo, photoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	// Pastikan user yang sedang login adalah pemilik foto
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userIDUint, _ := userID.(uint)
	if photo.UserID != userIDUint {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized to update this photo"})
		return
	}

	// Bind data yang diperbarui dari request
	if err := c.ShouldBindJSON(&photo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Simpan perubahan ke database
	if err := models.DB.Save(&photo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update photo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Photo updated successfully", "photo": photo})
}

// DeletePhotoHandler menghapus photo berdasarkan photo ID
func DeletePhotoHandler(c *gin.Context) {
	photoID := c.Param("photoId")

	// Ambil photo dari database berdasarkan photo ID
	var photo models.Photo
	if err := models.DB.First(&photo, photoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	// Pastikan user yang sedang login adalah pemilik foto
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userIDUint, _ := userID.(uint)
	if photo.UserID != userIDUint {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized to delete this photo"})
		return
	}

	// Hapus photo dari database
	if err := models.DB.Delete(&photo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete photo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Photo deleted successfully"})
}

// GetPhotosHandler mengambil daftar semua photos
func GetPhotosHandler(c *gin.Context) {
	var photos []models.Photo
	if err := models.DB.Find(&photos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch photos"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"photos": photos})
}
