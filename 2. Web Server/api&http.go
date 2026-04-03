package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// INI ADALAH REST API, yaitu Database yang bisa diakses melalui HTTP request,
type User struct{
	ID int `json:"id"` // ID pengguna, ditandai dengan tag JSON untuk serialisasi
	Name string `json:"name"` // Nama pengguna, ditandai dengan tag JSON untuk serialisasi
	Email string `json:"email"` // Email pengguna, ditandai dengan tag JSON untuk serialisasi
}


// ini adalah rest API yaitu Database yang bisa diakses melalui HTTP request,
//  biasanya digunakan untuk membuat aplikasi web atau mobile yang membutuhkan 
// backend untuk menyimpan dan mengelola data pengguna.

func dapatkanUser(w http.ResponseWriter, r *http.Request){ // Fungsi handler untuk menangani permintaan HTTP pada endpoint /products
	dataorang := []User{
		{
			ID:    1,
			Name:  "John Doe",
			Email: "john.doe@example.com",
		},
		{
			ID:    2,
			Name:  "Jane Smith",
			Email: "jane.smith@example.com",
		},
	}
	json.NewEncoder(w).Encode(dataorang) // Mengirim data dalam format JSON sebagai respons HTTP
}

func main(){
	http.HandleFunc("/products", dapatkanUser) // r adalah objek request yang berisi informasi tentang permintaan HTTP yang diterima,
	fmt.Println("Server berjalan di http://localhost:8080")
	http.ListenAndServe(":8080", nil) // Menjalankan server HTTP pada port 8080
}