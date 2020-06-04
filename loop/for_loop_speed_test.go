package loop

import "testing"

func BenchmarkFor(b *testing.B) {
	var a int
	for i := 0; i < b.N; i++ {
		a++
	}
}

func BenchmarkFor2(b *testing.B) {
	var a int
	for i := 0; i <= b.N; i++ {
		a++
	}
}
