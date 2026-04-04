package main

import "golang.org/x/crypto/bcrypt"

// ============== PASSWORD HELPER FUNCTIONS ==============

// HashPassword mengubah password menjadi hash menggunakan bcrypt
// Ini untuk keamanan - password tidak disimpan langsung di database
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// VerifyPassword membandingkan password yang diinput dengan hash di database
// Mengembalikan true jika password cocok, false jika tidak
func VerifyPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
