package main

import "fmt"

func main() {
	runApplication()
}

func logging() {
	fmt.Println("End")
}

func runApplication() {
	defer logging()
	fmt.Println("Run App")
	//logging()
}
