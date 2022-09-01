package main

import "fmt"

type Blacklist func(string) bool

func loginUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked ", name)
	} else {
		fmt.Println("Welcome ", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "novin"
	}

	loginUser("novin", blacklist)
	loginUser("novin", func(s string) bool {
		return s == "novin"
	})
}
