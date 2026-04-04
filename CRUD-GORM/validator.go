package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// ============== VALIDATOR ==============
var validate = validator.New()

func ValidateStruct(s interface{}) error {
	return validate.Struct(s)
}

// explaination:
// 1. Kita menggunakan package github.com/go-playground/validator/v10 untuk melakukan validasi struct.
// 2. Variabel validate adalah instance dari validator yang bisa digunakan untuk memvalidasi struct.
// 3. Fungsi ValidateStruct menerima sebuah interface{} yang bisa berupa struct apapun, dan memanggil method Struct() dari validator untuk melakukan validasi berdasarkan tag yang ada pada struct tersebut. Jika validasi gagal, fungsi ini akan mengembalikan error yang berisi detail tentang kesalahan validasi.

// cara menggunakan di handler:
// 1. Setelah melakukan binding data JSON ke struct, panggil fungsi ValidateStruct dengan struct tersebut sebagai argumen.
// 2. Jika fungsi mengembalikan error, berarti ada kesalahan validasi, dan kita bisa mengembalikan response error ke client dengan detail kesalahan tersebut.
// 3. Jika tidak ada error, berarti data valid dan kita bisa melanjutkan proses penyimpanan ke database atau operasi lainnya.

// error response struct untuk mengembalikan pesan error yang konsisten ke client
type ErrorResponse struct {
	Code int `json:"code"`
	Message string `json:"message"`
}
// fungsi untuk mengembalikan response error dengan format yang konsisten
func respondError(c *gin.Context, status int, message string){
	c.JSON(status, ErrorResponse{
		Code: status,
		Message: message,
	})
}
// fungsi respondError ini bisa digunakan di semua handler untuk 
// engembalikan error dengan format yang sama, 
// sehingga client bisa dengan mudah menangani error tersebut.