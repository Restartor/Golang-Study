package main

import "fmt"
func operator() {
// perhitungan
a := 10
b := 3
fmt.Println(a + b) // 13
fmt.Println(a - b) // 7
fmt.Println(a * b) // 30
fmt.Println(a / b) // 3 (pembagian bilangan bulat)
fmt.Println(a % b) // 1 (sisa pembagian)

// Perbandingan
fmt.Println(a > b) // true
fmt.Println(a == b) // false
fmt.Println(a != b) // true

// Logika
x := true
y := false

fmt.Println(x && y) // AND // false
fmt.Println(x || y) // OR // true
fmt.Println(!x)     // NOT  // false
}
