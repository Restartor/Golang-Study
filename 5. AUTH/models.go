package main

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// ============== MODELS ==============

// User model untuk menyimpan data pengguna
type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Email     string    `gorm:"uniqueIndex" json:"email"`
	Password  string    `json:"-"` // tidak ditampilkan dalam JSON response
	Name      string    `json:"name"`
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
