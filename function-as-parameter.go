package main

import (
	"fmt"
	"strings"
)

// type declaration -> tidak panjang panjang
type Filter func(string) string

func main() {
	checkedFilter("novin", upperName)
}

func checkedFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println(nameFiltered)
}

func upperName(name string) string {
	if name == strings.ToUpper(name) {
		return "Sudah Uppercase"
	} else {
		return "Belum Uppercase"
	}
}

func checkFilter(name string) string {
	if name == "anjing" {
		return "censored"
	} else {
		return name
	}
}
