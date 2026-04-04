package main

import "fmt"

func worker(ch chan string) { // Fungsi worker yang akan dijalankan sebagai goroutine
	ch <- "Hello from worker!" // Mengirim pesan ke channel

}

func gowork(){
	ch := make(chan string) // Membuat channel untuk komunikasi antar goroutine
	go worker(ch) // Menjalankan fungsi worker sebagai goroutine
	result := <-ch // Menerima pesan dari channel
	fmt.Println(result) // Output: Hello from worker!
}

// Penjelasan:
// goroutine adalah fungsi yang dapat berjalan secara bersamaan dengan fungsi lainnya. 
// Dalam contoh di atas, fungsi worker dijalankan sebagai goroutine 
// menggunakan kata kunci "go". 
// Channel digunakan untuk berkomunikasi antara goroutine, 
// di mana worker mengirim pesan ke channel dan gowork menerima pesan tersebut.