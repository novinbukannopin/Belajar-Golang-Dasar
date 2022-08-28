package main

import "fmt"

func main() {
	fmt.Println(sumAll(1, 2, 3, 4, 5))

	slice := []int{3, 4, 5, 6, 7}
	fmt.Println(sumAll(slice...))
}

func sumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}
