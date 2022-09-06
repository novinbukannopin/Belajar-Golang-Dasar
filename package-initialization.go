package main

import (
	"Belajar-Golang-Dasar/database"
	"fmt"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
