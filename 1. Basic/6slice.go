package main

import "fmt"

func slice(){
		numbers := []int{1, 2, 3, 4} // Membuat slice dengan elemen awal 1, 2, 3, 4
		numbers = append(numbers, 5) //

	fmt.Println(numbers)
}
// Output: [1 2 3 4 5] 