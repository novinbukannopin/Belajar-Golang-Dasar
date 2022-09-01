package main

import "fmt"

func main() {
	runApp(false)
}

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("erorr : ", message)
	}
	fmt.Println("Aplikasi Selesai")
}

func runApp(error bool) {
	fmt.Println("Aplikasi berjalan")

	defer endApp()
	if error {
		panic("Aplikasi Error")
	}
}
