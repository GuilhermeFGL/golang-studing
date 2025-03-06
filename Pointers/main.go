package main

import "fmt"

func main() {
	var va1 int = 1
	var va2 int = va1 // copy and pass value

	fmt.Println(va1, va2, " ")

	va1++
	fmt.Println(va1, va2, " ")

	var va3 int = 3
	var pointer *int = &va3                  // & pass memory reference; *int is the type saying that this is a pointer
	fmt.Println(va3, pointer, *pointer, " ") // * gets the value of the memory reference

	va3++
	fmt.Println(va3, pointer, *pointer, " ")

	*pointer++
	fmt.Println(va3, pointer, *pointer, " ")
}
