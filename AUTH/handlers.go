package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ============== HANDLERS ==============

// Register handler untuk registrasi user baru
func Register(c *gin.Context) {
	var request RegisterRequest

	// Binding dan validasi request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{
			"error":   "Data tidak valid",
			"details": err.Error(),
		})
		return
	}

	// Cek apakah email sudah terdaftar
	var existingUser User
	if err := DB.Where("email = ?", request.Email).First(&existingUser).Error; err == nil {
		c.JSON(400, gin.H{
			"error": "Email sudah terdaftar",
		})
		return
	}

	// Hash password
	hashedPassword, err := HashPassword(request.Password)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Gagal membuat user",
		})
		return
	}

	// Buat user baru
	newUser := User{
		Email:    request.Email,
		Password: hashedPassword,
		Name:     request.Name,
	}

	// Simpan ke database
	if err := DB.Create(&newUser).Error; err != nil {
		c.JSON(500, gin.H{
			"error": "Gagal menyimpan user ke database",
		})
		return
	}

	// Generate JWT token
	token, err := GenerateJWT(&newUser)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Gagal membuat token",
		})
		return
	}

	c.JSON(201, gin.H{
		"message":    "User berhasil terdaftar",
		"token":      token,
		"user_id":    newUser.ID,
		"email":      newUser.Email,
		"name":       newUser.Name,
		"created_at": newUser.CreatedAt,
	})
}

// Login handler untuk login user
func Login(c *gin.Context) {
	var request LoginRequest

	// Binding dan validasi request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{
			"error":   "Data tidak valid",
			"details": err.Error(),
		})
		return
	}

	// Cari user berdasarkan email
	var user User
	if err := DB.Where("email = ?", request.Email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(401, gin.H{
				"error": "Email atau password salah",
			})
		} else {
			c.JSON(500, gin.H{
				"error": "Gagal mencari user",
			})
		}
		return
	}

	// Verifikasi password
	if !VerifyPassword(user.Password, request.Password) {
		c.JSON(401, gin.H{
			"error": "Email atau password salah",
		})
		return
	}

	// Generate JWT token
	token, err := GenerateJWT(&user)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Gagal membuat token",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Login berhasil",
		"token":   token,
		"user_id": user.ID,
		"email":   user.Email,
		"name":    user.Name,
	})
}

// Logout handler - dalam praktik real, ini bisa menggunakan token blacklist
func Logout(c *gin.Context) {
	// Dapatkan token dari header
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(400, gin.H{
			"error": "Token tidak ditemukan",
		})
		return
	}

	// Dalam implementasi real, Anda bisa menambahkan token ke blacklist
	// atau menghapus token dari cache (seperti Redis)
	// Untuk contoh ini, kami hanya mengirim response sukses

	c.JSON(200, gin.H{
		"message": "Logout berhasil",
	})
}
