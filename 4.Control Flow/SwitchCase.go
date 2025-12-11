package main

import "fmt"

func mainKedua() {
	fmt.Print("Masukkan nomor hari (1-7): ")
	var day int
	fmt.Scanln(&day)
	switch day {
	case 1:
		fmt.Println("Senin")
	case 2:
		fmt.Println("Selasa")
	case 3:
		fmt.Println("Rabu")
	case 4:
		fmt.Println("Kamis")
	case 5:
		fmt.Println("Jumat")
	case 6:
		fmt.Println("Sabtu")
	case 7:
		fmt.Println("Minggu")
	default:
		fmt.Println("Hari tidak valid")
	}
}

// Penjelasan:
// Pada contoh di atas, kita menggunakan pernyataan switch-case untuk menentukan nama hari berdasarkan nilai variabel day.
// Setiap case memeriksa apakah nilai day sesuai dengan angka tertentu (1 hingga 7) dan mencetak nama hari yang sesuai.
// Jika nilai day tidak sesuai dengan salah satu case, maka blok default akan dieksekusi, mencetak "Hari tidak valid".
// Switch-case adalah cara yang efisien untuk menangani banyak kondisi berdasarkan satu variabel di Golang.
