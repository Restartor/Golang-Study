/*
================== PENJELASAN JWT AUTHENTICATION (DARI PEMULA) ==================

APA ITU JWT?
JWT (JSON Web Token) adalah sistem tanda tangan digital yang digunakan untuk memverifikasi
identitas user tanpa perlu menyimpan data session di server. Bayangkan seperti kartu ID yang
diberikan pemerintah - setelah user login, mereka mendapat "kartu digital" (token) yang
membuktikan mereka siapa.

BAGAIMANA CARA KERJANYA?

1. USER REGISTER (DAFTAR):
   - User memberikan email, password, dan nama ke server
   - Server menyimpan password dalam bentuk "kode rahasia" (hash) agar aman
   - Server simpan data user di database
   - Server memberikan JWT token kepada user sebagai "kartu"

2. USER LOGIN (MASUK):
   - User memberikan email dan password
   - Server cek apakah kombinasi email+password benar
   - Jika benar → server buat JWT token baru dan kirim ke user
   - Jika salah → tolak login

3. USER MENGAKSES DATA YANG TERLINDUNGI (PROTECTED):
   - User mengirim JWT token mereka di setiap request
   - Server verifikasi apakah token valid dan belum expired
   - Jika valid → user boleh akses data
   - Jika tidak valid → tolak request

4. USER LOGOUT (KELUAR):
   - User buang token mereka
   - Token tidak bisa dipakai lagi

STRUKTUR JWT TOKEN (3 BAGIAN):
- HEADER: Berisi tipe token (JWT) dan algoritma yang digunakan (HS256)
- PAYLOAD: Berisi data user (user_id, email, nama) dan waktu expired
- SIGNATURE: Tanda tangan digital untuk memverifikasi token asli atau palsu

CONTOH SEDERHANA:
1. Seseorang datang ke toko dan register → toko beri dia kartu member
2. Kartu itu berisi nama dia, nomor ID, dan tanda tangan toko
3. Saat belanja, dia tunjuk kartu itu ke kasir
4. Kasir verifikasi tanda tangan dan beri diskon
5. Jika kartu palsu atau kadaluarsa → tidak dapat diskon

KEAMANAN PASSWORD:
Kami gunakan bcrypt (password hashing) - password tidak disimpan langsung di database,
tapi di-ubah jadi "hash" (kode acak). Saat user login, kami ubah password input mereka
jadi hash dan bandingkan dengan hash di database. Jadi meskipun database bocor, password
asli tetap aman.

==============================================================================
*/

package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)
var DB *gorm.DB // Variabel global untuk database connection, pastikan sudah diinisialisasi di main.go

// ============== CONSTANTS ==============
const JWTSecretKey = "your-secret-key-change-this-in-production"

// ============== MODELS ==============

// User model untuk menyimpan data pengguna
type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Email    string `gorm:"uniqueIndex" json:"email"`
	Password string `json:"-"` // tidak ditampilkan dalam JSON response
	Name     string `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

// RegisterRequest untuk menerima data register dari client
type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Name     string `json:"name" binding:"required"`
}

// LoginRequest untuk menerima data login dari client
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// JWTClaims untuk menyimpan informasi dalam token JWT
type JWTClaims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	jwt.RegisteredClaims
}

// AuthResponse untuk mengirim response token
type AuthResponse struct {
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expires_at"`
	Message   string `json:"message"`
}

// ============== HELPER FUNCTIONS ==============

// HashPassword mengubah password menjadi hash menggunakan bcrypt
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// VerifyPassword membandingkan password dengan hash
func VerifyPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

// GenerateJWT membuat token JWT untuk user
func GenerateJWT(user *User) (string, error) {
	expiresAt := time.Now().Add(24 * time.Hour) // Token berlaku 24 jam

	claims := JWTClaims{
		UserID: user.ID,
		Email:  user.Email,
		Name:   user.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(JWTSecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateJWT memvalidasi token JWT
func ValidateJWT(tokenString string) (*JWTClaims, error) {
	claims := &JWTClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTSecretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("token tidak valid")
	}

	return claims, nil
}

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
		"message":  "User berhasil terdaftar",
		"token":    token,
		"user_id":  newUser.ID,
		"email":    newUser.Email,
		"name":     newUser.Name,
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
		"message":  "Login berhasil",
		"token":    token,
		"user_id":  user.ID,
		"email":    user.Email,
		"name":     user.Name,
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

// ============== MIDDLEWARE ==============

// AuthMiddleware middleware untuk memverifikasi JWT token
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

		// Simpan user info di context
		c.Set("user_id", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("name", claims.Name)

		c.Next()
	}
}

// ============== EXAMPLE PROTECTED ENDPOINT ==============

// GetProfile handler untuk mendapatkan profil user yang sedang login
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

// InitUserTable melakukan migrasi tabel User
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
