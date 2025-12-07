package main

import "fmt"

func main() {

	fmt.Println("Contoh berbagai gaya loop di golang: ")
	var decision int
	fmt.Scanln(&decision)

	switch decision {
	case 1:
		fmt.Println("forloopstyle")
		forloopstyle()
	case 2:
		fmt.Println("whileloopstyle")
		whileloopstyle()
	case 3:
		fmt.Println("infiniteLoopStyle")
		infiniteLoopStyle()
	default:
		fmt.Println("Gaya loop tidak valid")
	}

}

func forloopstyle() {
	// Loop standar
	for i := 1; i <= 5; i++ {
		fmt.Printf("Loop standar, iterasi ke-%d\n", i)
	}
}

func whileloopstyle() {
	// Loop gaya while
	count := 1
	for count <= 5 {
		fmt.Printf("Loop gaya while, iterasi ke-%d\n", count)
		count++
	}
}
func infiniteLoopStyle() {
	// Loop tak terbatas
	i := 1
	for {
		fmt.Printf("Loop tak terbatas, iterasi ke-%d\n", i)
		i++
	}
}
