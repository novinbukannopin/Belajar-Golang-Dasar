package main

import "fmt"

func main() {

	// manual
	var names [3]string
	names[0] = "novin"
	names[1] = "ardian"
	names[2] = "yulianto"

	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// auto declare
	var values = [3]int{
		90, 95, 80,
	}

	fmt.Println(values)
	fmt.Println(len(values))
	fmt.Println(len(names))
}
