package gopool

import "testing"

func BenchmarkSimplestPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SimplestPool()
	}
}

func BenchmarkNoPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NoPool()
	}
}
