package main

import "fmt"

func main() {
	n := 10

	if n > 15 {
		fmt.Println("bigger than 15")
	} else {
		fmt.Println("less than 15")
	}

	if m := n; m > 15 { // m is not visible outside this scope
		fmt.Println("still bigger than 15")
	} else {
		fmt.Println("still less than 15")
	}

	switch t := "opt1"; t {
	case "opt1":
		fmt.Println("opt1")
		fallthrough // continue to the next condition

	case "opt2":
		fmt.Println("opt2")

	default:
		fmt.Println("default")
	}
}
