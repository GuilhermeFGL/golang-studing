package main

import (
	"example.com/m/v2/AutomatedTests/addresses"
	"fmt"
)

func main() {
	address := "Street Main number 123"

	result := addresses.TypeOfAddress(address)
	fmt.Println(result)
}
