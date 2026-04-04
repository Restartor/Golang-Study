/*
================================================================================
                          AUTHENTICATION SYSTEM - COMPLETE GUIDE
================================================================================

✅ SEMUA ISSUES SUDAH DIPERBAIKI! Berikut penjelasan lengkap setiap file:

================================================================================
                               FILE STRUCTURE
================================================================================

📁 5. AUTH/
  ├── 📄 models.go          - Semua struct models (User, Request, Claims)
  ├── 📄 password.go        - Fungsi hash & verify password dengan bcrypt
  ├── 📄 jwt.go             - GenerateJWT & ValidateJWT functions
  ├── 📄 handlers.go        - Register, Login, Logout endpoint handlers
  ├── 📄 middleware.go      - AuthMiddleware & GetProfile functions
  ├── 📄 db.go              - Database initialization & global var
  ├── 📄 example_main.go    - Contoh implementasi di main.go
  └── 📄 GUIDE.go           - File ini (dokumentasi lengkap)

================================================================================
                            ALUR KONEKSI ANTAR FILE
================================================================================

models.go
├── User (struct)
├── RegisterRequest (struct)
├── LoginRequest (struct)
├── JWTClaims (struct)
└── AuthResponse (struct)
    ↓ digunakan oleh
    
password.go
├── HashPassword()      ← dari handlers.go (Register)
└── VerifyPassword()    ← dari handlers.go (Login)

jwt.go
├── GenerateJWT()       ← dari handlers.go (Register & Login)
├── ValidateJWT()       ← dari middleware.go (AuthMiddleware)
└── JWTSecretKey        ← constant untuk signing

db.go
├── var DB              ← global variable
├── InitDB()            ← panggil di main.go
├── GetDB()             ← helper function
└── InitUserTable()     ← di middleware.go

handlers.go
├── Register()          ← POST /register
├── Login()             ← POST /login
├── Logout()            ← POST /logout
└── menggunakan:
    - DB (dari db.go)
    - User model (dari models.go)
    - RegisterRequest model (dari models.go)
    - LoginRequest model (dari models.go)
    - HashPassword() (dari password.go)
    - VerifyPassword() (dari password.go)
    - GenerateJWT() (dari jwt.go)

middleware.go
├── AuthMiddleware()    ← protect routes yang perlu auth
├── GetProfile()        ← contoh protected endpoint
└── InitUserTable()     ← setup database
    menggunakan:
    - DB (dari db.go)
    - User model (dari models.go)
    - ValidateJWT() (dari jwt.go)

================================================================================
                              SETUP & PENGGUNAAN
================================================================================

STEP 1: Setup Environment Variables
   Buat file .env di root project:
   
   DB_HOST=localhost
   DB_USER=postgres
   DB_PASSWORD=your_password
   DB_NAME=your_database
   DB_PORT=5432
   DB_SSLMODE=disable

STEP 2: Install Dependencies
   go get github.com/golang-jwt/jwt/v5
   go get golang.org/x/crypto/bcrypt
   go get github.com/gin-gonic/gin
   go get gorm.io/driver/postgres
   go get gorm.io/gorm
   go get github.com/joho/godotenv

STEP 3: Buat main.go (di folder yang sama dengan file AUTH ini)
   func main() {
       godotenv.Load()
       InitDB()
       InitUserTable()
       
       router := gin.Default()
       
       // Public routes
       router.POST("/register", Register)
       router.POST("/login", Login)
       
       // Protected routes
       router.GET("/profile", AuthMiddleware(), GetProfile)
       router.POST("/logout", AuthMiddleware(), Logout)
       
       router.Run(":8080")
   }

STEP 4: Run
   go run .

================================================================================
                            ENDPOINT TESTING
================================================================================

TOOL: Gunakan Postman, curl, atau VS Code REST Client

1️⃣ REGISTER
   POST http://localhost:8080/register
   Body (JSON):
   {
     "email": "user@example.com",
     "password": "password123",
     "name": "John Doe"
   }
   
   Response:
   {
     "message": "User berhasil terdaftar",
     "token": "eyJhbGc...",
     "user_id": 1,
     "email": "user@example.com",
     "name": "John Doe"
   }

2️⃣ LOGIN
   POST http://localhost:8080/login
   Body (JSON):
   {
     "email": "user@example.com",
     "password": "password123"
   }
   
   Response:
   {
     "message": "Login berhasil",
     "token": "eyJhbGc...",
     "user_id": 1,
     "email": "user@example.com",
     "name": "John Doe"
   }
   
   ⚠️ SIMPAN TOKEN INI UNTUK STEP BERIKUTNYA!

3️⃣ GET PROFILE (Protected)
   GET http://localhost:8080/profile
   Headers:
   Authorization: eyJhbGc...
   
   Response:
   {
     "message": "Ini adalah data profil user",
     "user_id": 1,
     "email": "user@example.com",
     "name": "John Doe"
   }

4️⃣ LOGOUT (Protected)
   POST http://localhost:8080/logout
   Headers:
   Authorization: eyJhbGc...
   
   Response:
   {
     "message": "Logout berhasil"
   }

================================================================================
                            PENJELASAN SINGKAT
================================================================================

🔑 REGISTER
   - Terima email, password, name
   - Validasi format email & password minimum 6 chars
   - Cek apakah email sudah ada
   - Hash password dengan bcrypt
   - Simpan user ke database
   - Generate JWT token
   - Return token + user data

🔓 LOGIN
   - Terima email & password
   - Cari user di database
   - Verifikasi password dengan bcrypt
   - Jika cocok: generate JWT token & return
   - Jika tidak: return error

🛡️ PROTECTED ENDPOINT (dengan AuthMiddleware)
   - Client mengirim token di header "Authorization"
   - Middleware validasi token
   - Jika valid: set user info di context & lanjut ke handler
   - Jika invalid: return 401 unauthorized

🚪 LOGOUT
   - Terima token dari header
   - Return success message
   - (Client disarankan untuk delete token dari storage)

================================================================================
                         STRUKTUR & TANGGUNG JAWAB
================================================================================

models.go       ← Data models saja (struct)
password.go     ← Password security (hash & verify)
jwt.go          ← JWT token creation & validation
db.go           ← Database connection & initialization
handlers.go     ← Business logic (Register, Login, Logout)
middleware.go   ← Request middleware & protected endpoints
README.go       ← API documentation
example_main.go ← Contoh implementasi
GUIDE.go        ← File ini (setup & usage guide)

================================================================================

Sekarang semuanya sudah terhubung dengan baik! 🚀

Questions?
- Baca example_main.go untuk contoh implementasi
- Baca README.go untuk endpoint documentation
- Baca jwt.go dan password.go untuk memahami security
*/