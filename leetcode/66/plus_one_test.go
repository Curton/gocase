package _6

import (
	"fmt"
	"strconv"
	"testing"
)

func TestPlusOne(t *testing.T) {
	digits := []int{9, 9, 9, 9}
	fmt.Print(PlusOne(digits))
}

func TestPlusOne2(t *testing.T) {
	digits := []int{9, 9, 9, 9}
	fmt.Print(PlusOne2(digits))
}

func BenchmarkPlusOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := strconv.Itoa(i)
		v := make([]int, 16)
		for i2 := range s {
			v = append(v, int(i2-'0'))
			PlusOne(v)
		}
	}
}

func BenchmarkPlusOne2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := strconv.Itoa(i)
		v := make([]int, 16)
		for i2 := range s {
			v = append(v, int(i2-'0'))
			PlusOne2(v)
		}
	}
}
