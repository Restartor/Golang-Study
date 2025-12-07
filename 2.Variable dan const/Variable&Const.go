package main

import "fmt"

func main() {
	fmt.Println("This is about Variable & Constant")

	// deklarasi variabel normal
	var name string = "Rafi"
	const age int = 24
	// deklarasi variabel singkat
	country := "INDONESIA"
	president := "JOKOWI KALCER"
	// deklarasi variabel gabungan
	var (
		city, province string = "Jakarta", "DKI Jakarta"
	)

	// Multiple Shorthand Variable Declaration
	school, major := "University of Indonesia", "Computer Science"

	// Constant Declaration
	const pi float64 = 3.14
	// deklarasi konstant gabungan
	const (
		speedOfLight     = 299792458 // in meters per second
		gravitationalAcc = 9.81      // in m/s^2
	)

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println(country, "\n", president)
	fmt.Println("City:", city)
	fmt.Println("Province:", province)
	fmt.Println("School:", school)
	fmt.Println("Major:", major)
	fmt.Println("Value of Pi:", pi)
	fmt.Println("Speed of Light:", speedOfLight, "m/s")
	fmt.Println("Gravitational Acceleration:", gravitationalAcc, "m/s^2")
	printSpecialTypes()
}

// Types: string, int, float64, bool, byte, rune, complex64
// Penjelasan: string untuk teks, int untuk bilangan bulat, float64 untuk bilangan desimal,
// bool untuk nilai benar/salah, byte untuk data biner, rune untuk karakter Unicode,
// complex64 untuk bilangan kompleks.

// contoh byte dan rune dan complex64
var (
	sampleByte    byte      = 'A'    // byte adalah representasi dari satu byte (8 bit)
	sampleRune    rune      = '世'    // rune adalah representasi dari sebuah kode poin Unicode
	sampleComplex complex64 = 3 + 4i // complex64 adalah representasi dari sebuah bilangan kompleks
)

// contoh pengsingkatan fungsi seperti void di bahasa lain
func printSpecialTypes() {
	fmt.Println("Sample Byte:", sampleByte)
	fmt.Println("Sample Rune:", sampleRune)
	fmt.Println("Sample Complex Number:", sampleComplex)
}
