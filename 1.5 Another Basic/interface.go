package main

import "fmt"

// Interface Shape mendefinisikan kontrak untuk tipe-tipe yang memiliki metode Area

type Shape interface {
	Area() float64 
	// Mendefinisikan metode Area yang harus diimplementasikan oleh semua tipe yang memenuhi 
	// interface Shape
}

type Square struct {
	Side float64 // Menyimpan panjang sisi dari persegi
}
func (s Square) Area() float64 { // Implementasi metode Area untuk Square
	return s.Side * s.Side // Menghitung luas persegi dengan mengalikan sisi dengan dirinya sendiri
}

func interFace(){
	var s Shape = Square{Side: 4} // Membuat variabel s dengan tipe Shape dan menginisialisasi dengan objek Square
	fmt.Println("Area of square:", s.Area()) // Output: Area of square: 16
}