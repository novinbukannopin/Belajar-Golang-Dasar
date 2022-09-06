package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("Novin")
	data.PushBack("Ardian")
	data.PushBack("Yulianto")
	data.PushFront("Budi")

	fmt.Println(data.Front().Value)
	fmt.Println(data.Back().Value)

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	//	reverse

	for e := data.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}

	fmt.Println(data.Len())
}
