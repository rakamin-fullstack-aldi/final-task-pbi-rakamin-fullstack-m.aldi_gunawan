package handlers

import (
	"net/http"

	"final-task-pbi-rakamin-fullstack-m.aldi_gunawan/jwt"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware adalah middleware untuk memverifikasi token JWT
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil token dari header Authorization
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			// Jika tidak ada token, lanjutkan ke handler berikutnya (tidak diperlukan untuk register)
			c.Next()
			return
		}

		// Verifikasi token
		userID, err := jwt.VerifyToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Set user ID ke dalam context untuk digunakan di handler
		c.Set("user_id", userID)

		// Lanjutkan ke handler berikutnya
		c.Next()
	}
}
