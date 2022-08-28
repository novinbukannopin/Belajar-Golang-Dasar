package main

import "fmt"

func main() {

	name := "Novin"

	switch name {
	case "Novin":
		fmt.Println("Hello Novin")
	case "Joko":
		fmt.Println("Hello Joko")
	default:
		fmt.Println("Kenalan donk")
	}

	// SWITCH SHORT STATEMENT
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	// SWITCH TANPA KONDISI
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama sudah benar")
	case length < 1:
		fmt.Println("Nama terlalu pendek")
	default:
		fmt.Println("Tidak punya nama ya kamu !!")
	}
}
