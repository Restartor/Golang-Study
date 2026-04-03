package main

import "fmt"

func updateAge(age *int) { // Fungsi untuk mengupdate umur menggunakan pointer
	fmt.Printf("Masukkan umur terbaru anda : ")
	fmt.Scanln(age) // Mengupdate nilai age melalui pointer

}

func pointer(){
	age := 10
	fmt.Println("Sebelum update:", age) // Output: Sebelum update: 10

	updateAge(&age) // Mengirim alamat variabel age ke fungsi updateAge
	fmt.Println("Setelah update:", age) // Output: Setelah update: 25
}
// kenapa ga pake updateAge(*age) ?
// Karena kita ingin mengirim alamat variabel age ke fungsi updateAge, 
// bukan nilai yang ditunjuk oleh pointer.

// * untuk mendeklarasikan pointer, 
// & untuk mendapatkan alamat variabel, 
// dan * untuk mengakses nilai yang ditunjuk oleh pointer.
