package main

import (
	"fmt"
	"github.com/badoux/checkmail"
	"modulo/assist"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	assist.Write()

	fmt.Println(checkmail.ValidateFormat("123456"))
}
