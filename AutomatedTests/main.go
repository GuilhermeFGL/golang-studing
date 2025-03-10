package main

import (
	"example.com/m/v2/AutomatedTests/addresses"
	"example.com/m/v2/AutomatedTests/forms"
	"fmt"
)

func main() {
	address := "Street Main number 123"

	result := addresses.TypeOfAddress(address)
	fmt.Println(result)

	f1 := forms.Rectangle{Height: 5, Width: 5}
	f2 := forms.Circle{Radius: 5}
	forms.WriteArea(f1)
	forms.WriteArea(f2)
}
