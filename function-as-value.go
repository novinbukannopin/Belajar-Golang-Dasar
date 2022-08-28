package main

import "fmt"

func main() {
	sayGoodBye := getGoodBye
	result := sayGoodBye("Novin")
	fmt.Println(result)
}

func getGoodBye(name string) string {
	return "Good Bye " + name
}
