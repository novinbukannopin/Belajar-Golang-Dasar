package main

import "fmt"

func main() {

	var ujian = 80
	var absensi = 80

	var lulusUjian = ujian > 60
	var lulusAbsensi = absensi > 60

	var lulus = lulusUjian && lulusAbsensi
	fmt.Println(lulus)
}
