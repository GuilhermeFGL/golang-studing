package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	n := 0
	for n < 10 {
		fmt.Println(n)
		n++
	}

	texts := [3]string{"a1", "a2", "a3"}
	for i, t := range texts {
		fmt.Println(i, t)
	}

	for i, t := range "TEXT" {
		fmt.Println(i, string(t))
	}

	myMap := map[string]int{
		"property1": 1,
		"property2": 2,
	}
	for k, v := range myMap {
		fmt.Println(k, v)
	}

	for {
		fmt.Println("infinite loop")
		break
	}
}
