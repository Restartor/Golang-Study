package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

/*
STRUKTUR FILE AUTHENTICATION:
- jwt.go         → Penjelasan JWT + fungsi GenerateJWT & ValidateJWT
- models.go      → Semua struct (User, RegisterRequest, LoginRequest, JWTClaims)
- password.go    → HashPassword & VerifyPassword functions
- handlers.go    → Register, Login, Logout endpoints
- middleware.go  → AuthMiddleware, GetProfile, InitUserTable

CARA MENGGUNAKAN:
Daftarkan routes di main.go seperti ini:

	import auth "path/to/auth/package"

	func setupRoutes(router *gin.Engine) {
		// Public routes (tidak perlu token)
		router.POST("/register", auth.Register)
		router.POST("/login", auth.Login)

		// Protected routes (perlu token di header Authorization)
		router.GET("/profile", auth.AuthMiddleware(), auth.GetProfile)
		router.POST("/logout", auth.AuthMiddleware(), auth.Logout)
	}

JIKA PAKET MAIN DI FOLDER LAIN:
1. Copy seluruh folder "5. AUTH" ke project folder root
2. Update import path sesuai struktur project Anda
3. Pastikan variable global 'DB *gorm.DB' sudah diset sebelum menggunakan fungsi-fungsi auth
*/

// ExampleMainSetup adalah contoh bagaimana setup authentication di aplikasi Utama
func ExampleMainSetup() {
	// STEP 1: Connect ke database terlebih dahulu
	// ConnectDatabase() // dari connectDB.go

	// STEP 2: Initialize User table
	// InitUserTable()

	// STEP 3: Setup routes
	router := gin.Default()

	// Public routes
	router.POST("/register", Register)
	router.POST("/login", Login)

	// Protected routes
	router.GET("/profile", AuthMiddleware(), GetProfile)
	router.POST("/logout", AuthMiddleware(), Logout)

	// Run server
	log.Println("Server berjalan di :8080")
	router.Run(":8080")
}

/*
ENDPOINT EXAMPLES:

1. REGISTER (POST /register)
   Body: 
   {
     "email": "user@example.com",
     "password": "password123",
     "name": "John Doe"
   }
   
   Response (201):
   {
     "message": "User berhasil terdaftar",
     "token": "eyJhbGc...",
     "user_id": 1,
     "email": "user@example.com",
     "name": "John Doe",
     "created_at": "2026-04-04T10:30:00Z"
   }

2. LOGIN (POST /login)
   Body:
   {
     "email": "user@example.com",
     "password": "password123"
   }
   
   Response (200):
   {
     "message": "Login berhasil",
     "token": "eyJhbGc...",
     "user_id": 1,
     "email": "user@example.com",
     "name": "John Doe"
   }

3. GET PROFILE (GET /profile)
   Header: Authorization: eyJhbGc...
   
   Response (200):
   {
     "message": "Ini adalah data profil user",
     "user_id": 1,
     "email": "user@example.com",
     "name": "John Doe"
   }

4. LOGOUT (POST /logout)
   Header: Authorization: eyJhbGc...
   
   Response (200):
   {
     "message": "Logout berhasil"
   }
*/
