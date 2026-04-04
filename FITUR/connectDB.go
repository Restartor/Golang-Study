package main

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Mahasiswa struct {
	ID uint `gorm:"primaryKey"` // ID mahasiswa, ditandai sebagai primary key
	Nama string // Nama mahasiswa
	Umur int // Umur mahasiswa
}

var DB *gorm.DB // Variabel global untuk menyimpan koneksi database

func ConnectDatabase(){
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", // Menyusun string koneksi (DSN) untuk database menggunakan variabel lingkungan
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}) // Membuka koneksi ke database menggunakan GORM
	if err != nil{ log.Fatal("GAGAL KONEKSI KE DATABASE!") }
	db.AutoMigrate(&Mahasiswa{}) // Melakukan migrasi otomatis untuk membuat tabel Mahasiswa jika belum ada

	err = db.AutoMigrate(
		&Mahasiswa{},
	)
	if err != nil { log.Fatal("GAGAL MELAKUKAN MIGRASI!")}

	DB = db // Menyimpan koneksi database dalam variabel global DB untuk digunakan di seluruh aplikasi

	fmt.Println("Berhasil koneksi ke database")

}
