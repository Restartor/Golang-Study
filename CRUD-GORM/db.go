package main

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ============== DATABASE CONNECTION ==============
var DB *gorm.DB
func ConnectDB() {
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

	DB.AutoMigrate(
		&Product{},
	)	

	log.Println("Database berhasil terhubung")
}

// explaination:
// 1. Kita menggunakan package gorm.io/driver/postgres untuk driver PostgreSQL dan gorm.io/gorm untuk ORM.
// 2. Variabel global DB digunakan untuk menyimpan koneksi database yang bisa diakses di seluruh aplikasi.
// 3. Fungsi ConnectDB() membuat koneksi ke database menggunakan environment variables untuk konfigurasi.
// 4. Setelah koneksi berhasil, kita melakukan AutoMigrate untuk memastikan tabel Product sudah ada di database.
