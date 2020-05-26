package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInterface(t *testing.T) {
	var i interface{}
	i = "hello"
	fmt.Println(reflect.TypeOf(i))
	i = 1
	fmt.Println(reflect.TypeOf(i))
	i = 1.1
	fmt.Println(reflect.TypeOf(i))
	i = true
	fmt.Println(reflect.TypeOf(i))
}

// ha ha
func TestSum(t *testing.T) {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
}
