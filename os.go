package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println(args)

	name, err := os.Hostname()
	if err == nil {
		fmt.Println(name)
	} else {
		fmt.Println(err.Error())
	}
}
