package main

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB // Variabel global untuk database connection

// InitDB menginisialisasi koneksi database PostgreSQL
// Pastikan environment variables sudah diset:
// - DB_HOST
// - DB_USER
// - DB_PASSWORD
// - DB_NAME
// - DB_PORT
// - DB_SSLMODE
func InitDB() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal koneksi ke database:", err)
	}

	log.Println("Database berhasil terhubung")
}

// GetDB mengembalikan instance database
func GetDB() *gorm.DB {
	if DB == nil {
		log.Fatal("Database belum diinisialisasi. Panggil InitDB() terlebih dahulu!")
	}
	return DB
}
