package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var text = []byte("ashfdhkjadfbkjbvjfdsbkgbkdfjdsfbhajskdfghkabhlfd")

func TestAppend(t *testing.T) {
	assert.Equal(t, []byte("ashfdhkjadfbkjbvjfdsbkgbkdfjdsfbhajskdfghkabhlfd"), Append(text), "TestAppendDigitsErr")
}

func TestAppend2(t *testing.T) {
	assert.Equal(t, []byte("ashfdhkjadfbkjbvjfdsbkgbkdfjdsfbhajskdfghkabhlfd"), Append2(text), "TestAppendDigitsErr")
}

func TestCopy(t *testing.T) {
	assert.Equal(t, []byte("ashfdhkjadfbkjbvjfdsbkgbkdfjdsfbhajskdfghkabhlfd"), Copy(text), "CopyDigitsErr")
}

func BenchmarkAppendDigits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Append(text)
	}
}

func BenchmarkAppendDigits2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Append2(text)
	}
}

func BenchmarkCopyDigits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Copy(text)
	}
}

/*
goos: darwin
goarch: amd64
pkg: gocase/slice
BenchmarkAppendDigits-4         180540025               33.2 ns/op            48 B/op          1 allocs/op
BenchmarkAppendDigits2-4        210423793               28.4 ns/op            48 B/op          1 allocs/op
BenchmarkCopyDigits-4           212710212               28.3 ns/op            48 B/op          1 allocs/op
PASS
*/
