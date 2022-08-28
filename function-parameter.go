package main

import "fmt"

func main() {
	sayHelloTo("Novin", "Ardian")
}

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}
