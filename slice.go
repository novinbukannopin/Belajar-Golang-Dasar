package main

import "fmt"

func main() {
	var months = [...]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}

	// pointer	-> data pertama slice
	// len		-> len slice
	// capacity -> len pointer sampai len array
	fmt.Println(len(months[3:7]))
	fmt.Println(cap(months[3:7]))
	fmt.Println(months[3:7])

	// APPEND SLICE
	var slice2 = months[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Muharrom")
	// meskipun array out of bound, append bisa bikin array baru
	fmt.Println(slice3)
	slice3[1] = "Desember baru"
	fmt.Println(slice3)
	fmt.Println(slice2)
	fmt.Println(months)

	// MAKE SLICE
	newSlice := make([]string, 2, 5)
	newSlice[0] = "novin"
	newSlice[1] = "novan"
	fmt.Println(newSlice)

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	iniArray := [5]int{1, 2, 3, 4, 5}
	// ATAU [...]
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}
