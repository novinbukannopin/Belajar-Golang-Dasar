package main

import "fmt"

func main() {
	counter := 1

	for counter <= 5 {
		fmt.Println("Loop ke ", counter)
		counter++
	}

	//	for dengan statement
	for i := 0; i <= 5; i++ {
		fmt.Println("Loop ke", i)
	}

	slice := []string{"novin", "ardian", "yulianto"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, s := range slice {
		fmt.Println("Index", i, "=", s)
	}

	person := map[string]string{
		"name": "novin",
		"age":  "20",
	}

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
