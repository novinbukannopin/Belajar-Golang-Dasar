package main

import "fmt"

type Man struct {
	Name string
}

func main() {
	novin := Man{"Novin"}
	novin.Married()

	fmt.Println(novin.Name)
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}
