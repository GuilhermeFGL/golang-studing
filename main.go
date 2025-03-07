package main

import (
	"example.com/m/v2/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello, Go!")

	application := app.GenerateCommandLineApplication()
	err := application.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
