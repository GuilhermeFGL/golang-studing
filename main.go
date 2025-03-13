package main

import (
	"example.com/m/v2/App"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello, Go!")

	application := LocalApp.GenerateCommandLineApplication()
	err := application.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
