package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

// ============== MIDDLEWARE ==============

// AuthMiddleware middleware untuk memverifikasi JWT token
// Gunakan di routes yang memerlukan otentikasi
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil token dari header Authorization
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(401, gin.H{
				"error": "Token tidak ditemukan",
			})
			c.Abort()
			return
		}

		// Validasi token
		claims, err := ValidateJWT(token)
		if err != nil {
			c.JSON(401, gin.H{
				"error": "Token tidak valid atau sudah kadaluarsa",
			})
			c.Abort()
			return
		}

		// Simpan user info di context untuk digunakan di handler berikutnya
		c.Set("user_id", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("name", claims.Name)

		c.Next()
	}
}

// ============== EXAMPLE PROTECTED ENDPOINT ==============

// GetProfile handler untuk mendapatkan profil user yang sedang login
// Route ini harus di-protect dengan AuthMiddleware
func GetProfile(c *gin.Context) {
	// Ambil user info dari context (sudah tersimpan oleh middleware)
	userID := c.GetUint("user_id")
	email := c.GetString("email")
	name := c.GetString("name")

	c.JSON(200, gin.H{
		"message": "Ini adalah data profil user",
		"user_id": userID,
		"email":   email,
		"name":    name,
	})
}

// ============== SETUP DATABASE UNTUK USER ==============

// InitUserTable melakukan migrasi tabel User di database
// Panggil function ini di main() setelah ConnectDatabase()
func InitUserTable() {
	if DB == nil {
		log.Fatal("Database tidak diinisialisasi. Panggil ConnectDatabase() terlebih dahulu")
	}

	err := DB.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("Gagal melakukan migrasi User table:", err)
	}

	log.Println("User table sudah ready")
}

// JANGAN LUPA BUAT LIMITERATE UNTUK ENDPOINT LOGIN DAN REGISTRASI DI PRODUCTION UNTUK MENCEGAH BRUTE FORCE ATTACK!
// KITA BISA GUNAKAN PACKAGE SEPERTI github.com/ulule/limiter UNTUK IMPLEMENTASI RATE LIMITER DI GIN.

// untuk sementara kita gunakan golang.org/x/time/rate untuk implementasi sederhana di handler login dan register,
// tapi untuk implementasi yang lebih lengkap dan mudah digunakan, sebaiknya gunakan package limiter yang sudah ada.
