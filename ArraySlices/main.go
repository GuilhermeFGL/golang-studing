package main

import "fmt"

func main() {
	var array1 [5]string // empty array of 5 positions
	array1[0] = "a"
	// ...

	array2 := [5]string{"a"}

	array3 := [...]string{"a", "b", "c", "d", "e", "f"}

	array4 := array3[1:4] // points to memory getting positions 1, 2 and 3

	slice := []int{0, 1, 2} // non fixed size
	slice = append(slice, 3)

	fmt.Println(array1, array2, array3, array4, slice, " ")
}
