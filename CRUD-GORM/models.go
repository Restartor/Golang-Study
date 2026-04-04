package main

import (
	"gorm.io/gorm"
)

// ============== MODELS ==============

type Product struct {
    gorm.Model
    Name  string  `json:"name"`
    Price float64 `json:"price"`
    Stock int     `json:"stock"`
	Picture string  `json:"picture"`
	Description string  `json:"description"`
}

// explaination:
// 1. Struct Product mewakili model data untuk produk yang akan disimpan di database.
// 2. gorm.Model adalah struct bawaan GORM yang menyediakan field ID, CreatedAt, UpdatedAt, dan DeletedAt secara otomatis.
// 3. Field Name, Price, Stock, Picture, dan Description adalah atribut dari produk yang akan diisi saat membuat atau memperbarui data produk.
// 4. Tag json digunakan untuk menentukan nama field saat data dikirim atau diterima dalam format JSON melalui API.