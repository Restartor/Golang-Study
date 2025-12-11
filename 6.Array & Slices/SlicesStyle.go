package main

import "fmt"

func utama() {

	orangGanteng := []string{"jefri", "udin", "otong", "asep"}
	// membuat slice dengan cara singkat bisa integer atau string seperti diatas

	fmt.Println("Orang Ganteng awal:", orangGanteng)

	orangGanteng = append(orangGanteng, "dudung") // menambahkan elemen baru ke slice

	fmt.Println("Orang Ganteng:", orangGanteng)

	hAlo := "hello"
	hAlo = string(append([]rune(hAlo), '!'))
	fmt.Println(hAlo)

}

// Function tersebut menunjukkan cara menggunakan slice secara singkat
// di Golang tanpa mendefinisikan ukuran slice secara eksplisit.
