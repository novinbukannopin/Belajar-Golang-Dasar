package main

import (
	"Belajar-Golang-Dasar/helper"
	"fmt"
)

func main() {
	helper.SayHello("Novin")
	fmt.Println(helper.Application)
	// tidak bisa
	//fmt.Println(helper.version)
}
