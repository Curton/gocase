package main

import "fmt"

func main() {
	var s string
	s = "some words"
	b := []byte(s)
	fmt.Printf("%s", b)
}
