package main

import "fmt"

func main() {
	var name string
	name = "Novin Ardian Yulianto"

	fmt.Println(name)

	//replace
	name = "Novin Ganteng"
	fmt.Println(name)

	// create var without type data
	var alias = "Novin"
	fmt.Println(alias)

	var age = 20
	fmt.Println(age)

	// tidak wajib use var

	nickname := "Novin"
	fmt.Println(nickname)

	// create multiple var
	var (
		firstName = "Novin"
		lastName  = "Ardian"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
}
