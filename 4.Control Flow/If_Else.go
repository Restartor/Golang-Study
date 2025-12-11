package main

import "fmt"

func main() {
	age := 20
	if age > 18 {
		fmt.Println("Kamu sudah dewasa")
	} else {
		fmt.Println("Kamu belum dewasa")
	}

}

// Penjelasan:
// Pada contoh di atas, kita menggunakan pernyataan if-else untuk memeriksa apakah variabel age lebih besar dari 18.
// Jika kondisi terpenuhi (age > 18), maka blok kode di dalam if akan dieksekusi, mencetak "Kamu sudah dewasa".
// Jika kondisi tidak terpenuhi, maka blok kode di dalam else akan dieksekusi, mencetak "Kamu belum dewasa".
// Ini adalah cara dasar untuk mengontrol alur program berdasarkan kondisi tertentu di Golang.
