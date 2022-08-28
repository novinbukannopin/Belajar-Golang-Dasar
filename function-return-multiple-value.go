package main

import "fmt"

func main() {
	// _ untuk ignore paramater yang tidak butuh
	firstName, _ := getFullName("Novin", "Ardian")
	fmt.Println(firstName)
}
func getFullName(firstName string, lastName string) (string, string) {
	return firstName, lastName
}
