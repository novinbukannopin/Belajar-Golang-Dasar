package helper

import "fmt"

// ini diawali huruf kapital = can
var Application = "APP"

// ini diawali huruf kecil = cant
var version = 1

func SayHello(name string) {
	fmt.Println("Hai", name)

	//	fn atau var yang diawali huruf kecil,
	//	dia tidak bisa di modifier keluar dari packagenya

}
