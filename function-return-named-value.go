package main

import "fmt"

func main() {
	firstName, middleName, lastName := getCompleteName()
	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
}
func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Novin"
	middleName = "Ardian"
	lastName = "Yulianto"
	return
}
