package main

import "fmt"

func main() {
	fmt.Println("This is about input dan output di Golang")
	fmt.Print("Masukkan nama Anda: ")
	var name string
	fmt.Scanln(&name)
	fmt.Printf("hallo %s Selamat datang di Golang\n", name)

	fmt.Print("Masukkan umur Anda: ")
	var age int
	fmt.Scanln(&age)
	fmt.Printf("Umur anda adalah %d tahun\n", age)
}

// Penjelasan:
// Untuk input dan output di Golang.
// Println adalah untuk mencetak output dengan baris baru di akhir.
// Printf adalah untuk mencetak output dengan format tertentu.
// Print adalah untuk mencetak output tanpa baris baru di akhir.
// Scanln adalah untuk membaca input dari pengguna hingga baris baru.
// & digunakan untuk mendapatkan alamat variabel yang akan diisi input dari pengguna.
// Contoh di atas meminta pengguna untuk memasukkan nama dan umur, lalu mencetaknya kembali ke layar.

// %s dan %d adalah format spesifier di Printf.
// %s digunakan untuk string.
// %d digunakan untuk integer.
