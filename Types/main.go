package main

import (
	"errors"
	"fmt"
)

func main() {

	// numbers
	var n int = 100000 //int8, int32, unint ...
	var b byte = 1
	var f float32 = 1 // float64

	// text
	var s string = "hello world"
	var c = 'c' // char

	// boolean
	var bol bool = true

	// error
	var e error = errors.New("internal error")

	fmt.Println(n, b, f, s, c, bol, e, " ")
}
