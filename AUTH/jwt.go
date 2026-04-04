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
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// ============== CONSTANTS ==============
const JWTSecretKey = "your-secret-key-change-this-in-production"

// ============== JWT CORE FUNCTIONS ==============

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
