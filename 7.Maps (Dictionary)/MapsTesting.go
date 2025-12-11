package main

import "fmt"

func main() {

	//pembuatan map
	kelas := map[string]int{
		"jefri": 90,
		"udin":  80,
		"otong": 70,
		"asep":  60,
	}

	fmt.Println(kelas)
	fmt.Println("==========================")

	// memisahkan nilai dari map
	fmt.Println("Nilai jefri:", kelas["jefri"])
	fmt.Println("Nilai udin:", kelas["udin"])

	// Mengubah nilai dari map
	fmt.Println("==========================")
	kelas["otong"] = 75
	fmt.Println("Nilai otong setelah diubah:", kelas["otong"])

	// Menambahkan elemen baru ke map
	fmt.Println("==========================")

	kelas["anjas"] = 85
	fmt.Println(kelas)
	fmt.Println("Murid baru Anjas dengan nilai: ", kelas["anjas"])

	// Menghapus elemen dari map
	fmt.Println("==========================")
	delete(kelas, "asep")
	fmt.Println("Setelah menghapus Asep:", kelas)

	// Mengecek keberadaan kunci dalam map
	fmt.Println("==========================")
	nilai, ada := kelas["asep"]
	if ada {
		fmt.Println("Nilai Asep:", nilai)
	} else {
		fmt.Println("Asep tidak ditemukan dalam kelas.")
	}

	// Iterasi map dengan for range
	fmt.Println("==========================")
	for nama, nilai := range kelas {
		fmt.Printf("Nilai %s: %d\n", nama, nilai)
	}

	// Membuat map kosong dengan make
	fmt.Println("==========================")
	kelasBaru := make(map[string]int)
	kelasBaru["budi"] = 88
	kelasBaru["siti"] = 92
	fmt.Println("Map kelas baru:", kelasBaru)

	// panjang map
	fmt.Println("==========================")
	fmt.Println("Jumlah item di kelas : ", len(kelas))
	fmt.Println("Jumlah item di kelasBaru : ", len(kelasBaru))

}
