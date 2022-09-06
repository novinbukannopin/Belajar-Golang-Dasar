package main

import "fmt"

// ini pass by value != java atau php
// karena dia me replace data, menyimpan di memori addres baru
func main() {
	address1 := Address{"Lamongan", "Jawa Timur", "Indonesia"}
	address2 := address1

	address2.City = "Gresik"
	fmt.Println(address1)
	fmt.Println(address2)

	// kalau pass by reference, dia perlu &
	// ini default di java
	address3 := Address{"Surabaya", "Jawa Timur", "Indonesia"}
	address4 := &address3

	address3.City = "Nganjuk"

	// tanda * berarti dia pointer ke -> tujuan
	fmt.Println(address3)
	fmt.Println(address4)

	alamat1 := new(Address)
	alamat1.City = "Surabaya"
	fmt.Println(alamat1)

	alamat := Address{"Surabaya", "Jawa Timur", ""}
	// dia tujuan pointer
	changeCountryToIndonesia(&alamat)
	fmt.Println(alamat)

}

// dia source pointer
func changeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

type Address struct {
	City, Province, Country string
}
