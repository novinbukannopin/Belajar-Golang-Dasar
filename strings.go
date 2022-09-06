package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Novin Ardian Yulianto", "Novin"))
	fmt.Println(strings.Split("Novin Ardian Yulianto", " "))
	fmt.Println(strings.Trim("    Novin Ardian Yulianto    ", " "))
	fmt.Println(strings.ReplaceAll("Novin Ardian Yulianto", "Novin", "Uhuy"))
}
