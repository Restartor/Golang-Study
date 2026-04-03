package main

import "fmt"

func inputoutput(){
	fmt.Println("hello world")
	var name string
	fmt.Print("Enter your name : ")
	fmt.Scanln(&name)
	fmt.Println("Hello,", name)
}

