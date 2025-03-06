package main

import "fmt"

func main() {
	var array1 [5]string // empty array of 5 positions
	array1[0] = "a"

	array2 := [5]string{"a"}

	array3 := [...]string{"a", "b", "c", "d", "e", "f"}

	array4 := array3[1:4] // points to memory getting positions 1, 2 and 3

	slice1 := []int{0, 1, 2} // non fixed size
	slice1 = append(slice1, 3)

	slice2 := make([]int, 10, 15) // type, initial number of items, max capacity that when reached it will be doubled

	fmt.Println(array1, array2, array3, array4, slice1, slice2, len(slice2), cap(slice2), " ")
}
