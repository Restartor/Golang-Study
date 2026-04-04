package main

import (
	"gorm.io/gorm"
)

// ============== MODELS ==============

type Product struct {
    gorm.Model
    Name  string  `json:"name" validate:"required"`
    Price float64 `json:"price" validate:"required,gt=0"`
    Stock int     `json:"stock" validate:"required,gt=0"`
	Picture string  `json:"picture" validate:"required,url"`
	Description string  `json:"description" validate:"required"`
}

// explaination:
// 1. Struct Product mewakili model data untuk produk yang akan disimpan di database.
// 2. gorm.Model adalah struct bawaan GORM yang menyediakan field ID, CreatedAt, UpdatedAt, dan DeletedAt secara otomatis.
// 3. Field Name, Price, Stock, Picture, dan Description adalah atribut dari produk yang akan diisi saat membuat atau memperbarui data produk.
// 4. Tag json digunakan untuk menentukan nama field saat data dikirim atau diterima dalam format JSON melalui API.
// 5. Dengan menggunakan GORM, kita bisa dengan mudah melakukan operasi CRUD pada model Product tanpa harus menulis query SQL secara manual.


// jangan lupa untuk validate input data di handler sebelum menyimpan ke database, 
// agar tidak di bruteforce dengan data yang tidak valid atau berbahaya.

// gt=0 artinya nilai harus lebih besar dari 0, ini untuk memastikan harga dan stok tidak negatif atau nol.
// url artinya nilai harus berupa URL yang valid, ini untuk memastikan field Picture berisi link gambar yang benar.
// required artinya field tersebut wajib diisi, ini untuk memastikan semua data penting produk terisi dengan benar.