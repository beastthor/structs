package main

import "fmt"

type contatctInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contatctInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contatctInfo: contatctInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	fmt.Printf("%+v", jim)

}
