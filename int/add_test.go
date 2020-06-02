package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func BenchmarkAdd01(b *testing.B) {
	var a = 0
	for i := 0; i < b.N; i++ {
		a += 1
	}
	assert.Equal(b, a, b.N, "Err in BenchmarkAdd01")
}

func BenchmarkAdd02(b *testing.B) {
	var a = 0
	for i := 0; i < b.N; i++ {
		a -= -1
	}
	assert.Equal(b, a, b.N, "Err in BenchmarkAdd02")
}

func BenchmarkAdd03(b *testing.B) {
	var a = 0
	for i := 0; i < b.N; i++ {
		a = a + 1
	}
	assert.Equal(b, a, b.N, "Err in BenchmarkAdd03")
}

func BenchmarkAdd04(b *testing.B) {
	var a = 0
	for i := 0; i < b.N; i++ {
		a++
	}
	assert.Equal(b, a, b.N, "Err in BenchmarkAdd04")
}

//goos: darwin
//goarch: amd64
//pkg: gocase/int
//BenchmarkAdd01
//BenchmarkAdd01-4   	1000000000	         0.515 ns/op
//BenchmarkAdd02
//BenchmarkAdd02-4   	1000000000	         0.261 ns/op
//BenchmarkAdd03
//BenchmarkAdd03-4   	1000000000	         0.514 ns/op
//BenchmarkAdd04
//BenchmarkAdd04-4   	1000000000	         0.258 ns/op
//PASS
