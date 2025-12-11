package main

import "fmt"

func main() {
	var pacar [4]string = [4]string{"kim minjeong", "haerin", "karina", "cha eunwoo"}

	fmt.Println("Array:", pacar)
	fmt.Println("Pacar pertama:", pacar[0])
	fmt.Println("Pacar kedua:", pacar[1])
	fmt.Println("Pacar ketiga:", pacar[2])
	fmt.Println("Pacar keempat:", pacar[3])

}

// contoh array normal di golang adalah diatas ini

func arraySingkat() {
	pacar := []string{"kim minjeong", "haerin", "karina", "cha eunwoo"}
	fmt.Println("Array singkat:", pacar)
}

// Function tersebut menunjukkan cara menggunakan array secara singkat
// di Golang tanpa mendefinisikan ukuran array secara eksplisit.
