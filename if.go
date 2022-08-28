package main

import "fmt"

func main() {

	var name = "eko"

	if name == "novin" {
		fmt.Println("Hello Novin")
	} else if name == "eko" {
		fmt.Println("Hello Eko")
	} else {
		fmt.Println("Hi, boleh kenalan?")
	}

	// IF SHORT STATEMENT
	if length := len(name); length > 5 {
		fmt.Println("Terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")

	}
}
