package main

import "fmt"

func add(n1 int, n2 int) int {
	return n1 + n2
}

func switchNumbers(n1, n2 int) (int, int) {
	return n2, n1
}

func mappedReturn(n1, n2 int) (r1 int, r2 int) {
	r1 = n2
	r2 = n1
	return
}

func main() {
	fmt.Println(add(1, 2))

	fmt.Println(switchNumbers(1, 2))

	r1, r2 := mappedReturn(1, 2)
	fmt.Println(r1, r2)

	internalFunction := func(text string) {
		fmt.Println(text)
	}
	internalFunction("hello world from a function")
}
