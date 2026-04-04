package main

/*
CONTOH IMPLEMENTASI:

File ini adalah contoh bagaimana menggunakan seluruh authentication system.
Jika Anda ingin menggunakannya, buat file main.go sendiri di folder yang sama
dengan semua file AUTH ini, kemudian gunakan struktur seperti contoh di bawah.

SETUP AWAL:
1. Pastikan semua file AUTH sudah ada di folder yang sama dengan main.go
2. Setup .env file dengan environment variables untuk database:
   DB_HOST=localhost
   DB_USER=postgres
   DB_PASSWORD=your_password
   DB_NAME=your_database
   DB_PORT=5432
   DB_SSLMODE=disable

3. Jalankan: go run .
*/

// CONTOH IMPLEMENTASI (aktifkan dengan uncomment dan hapus comment di func main)
/*
func main() {
	// Load environment variables dari .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file tidak ditemukan, menggunakan system environment variables")
	}

	// STEP 1: Initialize database
	InitDB()

	// STEP 2: Initialize User table (buat table jika belum ada)
	InitUserTable()

	// STEP 3: Setup Gin router
	router := gin.Default()

	// STEP 4: Register routes
	setupRoutes(router)

	// STEP 5: Start server
	log.Println("✅ Server berjalan di http://localhost:8080")
	router.Run(":8080")
}

// setupRoutes mengatur semua endpoint untuk authentication
func setupRoutes(router *gin.Engine) {
	// Health check
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Auth server berjalan dengan baik!",
		})
	})

	// === PUBLIC ROUTES (tidak perlu token) ===
	router.POST("/register", Register)
	router.POST("/login", Login)

	// === PROTECTED ROUTES (perlu JWT token di header Authorization) ===
	router.GET("/profile", AuthMiddleware(), GetProfile)
	router.POST("/logout", AuthMiddleware(), Logout)
}

// ============== TESTING ENDPOINTS ==============
/*
GUNAKAN TOOLS SEPERTI POSTMAN ATAU CURL UNTUK TEST:

1. REGISTER (POST http://localhost:8080/register)
   Body (JSON):
   {
     "email": "user@example.com",
     "password": "password123",
     "name": "John Doe"
   }

2. LOGIN (POST http://localhost:8080/login)
   Body (JSON):
   {
     "email": "user@example.com",
     "password": "password123"
   }

   Copy token dari response!

3. GET PROFILE (GET http://localhost:8080/profile)
   Headers:
   Authorization: <paste-token-dari-login>

4. LOGOUT (POST http://localhost:8080/logout)
   Headers:
   Authorization: <paste-token-dari-login>
*/
