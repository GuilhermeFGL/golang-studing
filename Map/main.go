package main

import "fmt"

func main() {
	myMap1 := map[string]int{
		"property1": 1,
		"property2": 2,
	}

	myMap2 := map[string]map[string]int{
		"property1": {
			"interProperty1": 1,
			"interProperty2": 2,
		},
	}
	myMap2["property2"] = map[string]int{
		"interProperty1": 1,
		"interProperty2": 2,
	}

	fmt.Println(myMap1, myMap2, " ")
}
