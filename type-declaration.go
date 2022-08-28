package main

import "fmt"

func main() {
	type NoKTP string
	type hasMarried bool

	var noKtpNovin NoKTP = "352426070903"
	var marriedStatus hasMarried = false
	fmt.Println(noKtpNovin)
	fmt.Println(marriedStatus)

}
