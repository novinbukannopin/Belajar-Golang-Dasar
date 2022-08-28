package main

import "fmt"

func main() {

	person := map[string]string{
		"name":    "novin",
		"address": "lamongan",
	}

	// add key
	person["title"] = "programmer"
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])

	var book = make(map[string]string)
	book["title"] = "Belajar GoLang"
	book["author"] = "Novin Ardian"
	book["ups"] = "Salah"

	fmt.Println(book)
	delete(book, "ups")
	fmt.Println(book)

}
