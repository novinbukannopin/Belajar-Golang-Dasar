package main

import "fmt"

func main() {
	// const wajib di declared by value
	const firstName string = "Novin"
	const lastName = "Ardian"

	fmt.Println(firstName)
	fmt.Println(lastName)

	// const cannot assign
	// tidak boleh diubah
	// constant tidak wajib di used

	// MULTIPLE CONSTANT
	const (
		country = "indonesia"
		addres  = "lamongan"
	)

	fmt.Println(country)
	fmt.Println(addres)
}
