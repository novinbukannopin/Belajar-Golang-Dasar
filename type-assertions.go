package main

import "fmt"

func main() {
	var result interface{} = random()
	// ini type assertions versi paksa
	//var resultString string = result.(string)
	//fmt.Println(resultString)
	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is int")
	default:
		fmt.Println("Unknown type")
	}
}

// kalo interface beda tipe dengan assertions
// dia akan panic

func random() interface{} {
	return "hai"
}
