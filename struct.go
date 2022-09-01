package main

import "fmt"

type Customer struct {
	name, address string
	age           int
}

func (customer Customer) sayHai(name string) {
	fmt.Println("Hello " + name + ", My Name is " + customer.name)
}

func main() {
	var novin Customer
	novin.name = "Novin "
	novin.address = "Indonesia"
	novin.age = 18
	fmt.Println(novin)

	ucup := Customer{
		name:    "ucup",
		address: "nganjuk",
		age:     90,
	}
	fmt.Println(ucup)

	udin := Customer{"udin", "zimbabwe", 17}
	fmt.Println(udin)

	novin.sayHai("Ucup")
}
