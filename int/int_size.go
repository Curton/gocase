package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a int
	b := 1
	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(b))

}
