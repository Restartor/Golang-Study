/*
STRUKTUR FILE AUTHENTICATION:
- jwt.go         → Penjelasan JWT + fungsi GenerateJWT & ValidateJWT
- models.go      → Semua struct (User, RegisterRequest, LoginRequest, JWTClaims)
- password.go    → HashPassword & VerifyPassword functions
- handlers.go    → Register, Login, Logout endpoints
- middleware.go  → AuthMiddleware, GetProfile, InitUserTable
- db.go          → Database connection dan initialization

CARA MENGGUNAKAN:
Daftarkan routes di main.go seperti ini:

	func setupRoutes(router *gin.Engine) {
		// Public routes (tidak perlu token)
		router.POST("/register", Register)
		router.POST("/login", Login)

		// Protected routes (perlu token di header Authorization)
		router.GET("/profile", AuthMiddleware(), GetProfile)
		router.POST("/logout", AuthMiddleware(), Logout)
	}

	func main() {
		ConnectDatabase() // Koneksi database
		InitUserTable()   // Buat/update tabel User
	
		router := gin.Default()
		setupRoutes(router)
		router.Run(":8080")
	}

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
     "name": "John Doe"
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
     "email": "user@example.com"
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
