package main

import (
	"fmt"
	"time"
)

func main() {
	go printOnConsole("concurrency 1") // go routine
	printOnConsole("concurrency 2")
}

func printOnConsole(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
