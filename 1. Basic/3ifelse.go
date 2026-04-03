package main

import "fmt"

func ifelse()  {
	
	nilai := 85
		if nilai >= 90 {
			fmt.Println("A")
		} else if nilai >= 75 {
			fmt.Println("B")
		} else {
			fmt.Println("C")
		}
}
// Output: B (karena nilai 85 lebih besar atau sama dengan 75 tetapi kurang dari 90)