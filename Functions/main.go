package main

import "fmt"

func add(n1 int, n2 int) int {
	return n1 + n2
}

func switchNumbers(n1, n2 int) (int, int) {
	return n2, n1
}

func mappedReturn(n1, n2 int) (r1 int, r2 int) {
	r1 = n2
	r2 = n1
	return
}

func multipleParams(numbers ...int) (total int) {
	for number := range numbers {
		total += number
	}
	return
}

func recursive(n int) int {
	if n > 1 {
		return recursive(n-2) + recursive(n-1)
	}
	return n
}

func delayedExecution() {
	defer fmt.Println("this will also be delayed")
	fmt.Println("delayed execution")
}

func panicAndRecover(i int) bool {
	defer func() {
		if r := recover(); r != nil { // recover from panic
			fmt.Println("Recovering from panic:", r)
		}
	}()

	if i > 0 {
		return true
	} else if i < 0 {
		return false
	}

	panic("!!! i is 0 !!!") // interrupt execution (this is not an error)
}

func closure() func() {
	closure := func() {
		fmt.Println("closure")
	}

	return closure
}

func pointer(n *int) {
	*n *= -1
}

func init() { // will run before main function; can have one init function per file
	fmt.Println("init")
}

func main() { // can have one main function per package
	defer delayedExecution() // will be executed at the end of the scope

	fmt.Println(add(1, 2))

	fmt.Println(switchNumbers(1, 2))

	r1, r2 := mappedReturn(1, 2)
	fmt.Println(r1, r2)

	fmt.Println(multipleParams(1, 2, 3, 4, 5, 6, 7, 7, 9))

	internalFunction := func(text string) {
		fmt.Println(text)
	}
	internalFunction("hello world from a function")

	fmt.Println("fibonacci:", recursive(10))

	panicAndRecover(0)

	closure()()
	myClosure := closure()
	myClosure()

	reference := 10
	pointer(&reference)
	fmt.Println(reference)

	result := func(text string) string {
		println("Anonymous function:", text) // will not be printed
		return text
	}("executing") // declare and execute the function
	fmt.Println(result)
}
