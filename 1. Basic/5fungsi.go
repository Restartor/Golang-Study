package main

import "fmt"

func tambah(a int, b int) int { // membuat fungi tambah a+b
    return a + b
}

func fungsi() { 
    hasil := tambah(5, 3) // Memanggil fungsi tambah
    fmt.Println(hasil) // Output: 8
}