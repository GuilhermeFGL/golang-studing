package main

import "fmt"

type user struct {
	name    string
	age     uint8
	address // this is "inheritance"
	// it will basically copy all it properties into this struct
}

type address struct {
	street string
	number int
}

func main() {
	var myUser1 = user{
		name: "Jack",
		age:  23,
		address: address{
			street: "123",
			number: 201,
		},
	}

	var myUser2 user
	myUser2.name = "David"
	myUser2.age = 21
	myUser2.address = address{}
	myUser2.address.street = "Avenue A"
	myUser2.address.number = 500

	myUser3 := user{"Ana", 19, address{"street", 404}}

	var myUser4 user
	var myUser5 = user{}

	fmt.Println(myUser1, myUser2, myUser3, myUser4, myUser5, " ")
}
