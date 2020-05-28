package main

import "fmt"

func main() {
	var s string
	s = "some words"
	b := []byte(s)
	fmt.Printf("%s \r\n", b)

	str := "你好,世界"
	fmt.Printf("%v \r\n", len(str))

	str2 := `a
b`
	fmt.Printf("%v \r\n", len(str2))
	fmt.Printf("%v", str2[1])
}
