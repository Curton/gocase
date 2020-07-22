/**
 * @author Covey Liu, covey@liukedun.com
 * @date 2020/7/13 14:54
 */

package main

import (
	"fmt"
	"testing"
)

func Benchmark(b *testing.B) {
	var k int
	for i := 0; i < b.N; i++ {
		j := int64(i)
		k += int(j)
	}
	fmt.Println(k)
}

func Benchmark2(b *testing.B) {
	var k int
	for i := 0; i < b.N; i++ {
		j := (int64)(i)
		k += (int)(j)
	}
	fmt.Println(k)
}
