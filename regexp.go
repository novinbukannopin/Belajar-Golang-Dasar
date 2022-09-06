package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("n([a-z])n")
	fmt.Println(regex.MatchString("adaa"))
}
