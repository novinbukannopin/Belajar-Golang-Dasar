package main

import (
	"flag"
	"fmt"
)

func main() {
	var host *string = flag.String("host", "localhost", "Put your host")
	var user *string = flag.String("user", "root", "Put your user")
	var password *string = flag.String("pw", "root", "Put your password")

	flag.Parse()
	fmt.Println("Host : ", *host)
	fmt.Println("Host : ", *user)
	fmt.Println("Host : ", *password)
}
