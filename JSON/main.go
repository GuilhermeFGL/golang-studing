package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type dog struct {
	Name string `json:"name"`
	Race string `json:"race"`
}

func main() {
	d1 := dog{"Ozzy", "Beagle"}
	json1, error1 := json.Marshal(d1)
	if error1 != nil {
		log.Fatal(error1)
	}
	fmt.Println(bytes.NewBuffer(json1))

	d2 := map[string]string{
		"name": "Bond",
		"race": "Ozzy",
	}
	json2, error2 := json.Marshal(d2)
	if error2 != nil {
		log.Fatal(error2)
	}
	fmt.Println(bytes.NewBuffer(json2))

	json3 := `{"name": "Ozzy", "race": "Beagle"}`
	var d3 dog
	if error3 := json.Unmarshal([]byte(json3), &d3); error3 != nil {
		log.Fatal(error3)
	}
	fmt.Println(d3)

	json4 := `{"name": "Bond", "race": "Beagle"}`
	d4 := make(map[string]string)
	if error4 := json.Unmarshal([]byte(json4), &d4); error4 != nil {
		log.Fatal(error4)
	}
	fmt.Println(d4)
}
