package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var text = []byte("ashfdhkjadfbkjbvjfdsbkgbkdfj88327598dsfbhajskdfghkabhlfd")

func TestAppendDigits(t *testing.T) {
	assert.Equal(t, []byte("88327598"), AppendDigits(text), "TestAppendDigitsErr")
}

func TestCopyDigits(t *testing.T) {
	assert.Equal(t, []byte("88327598"), CopyDigits(text), "CopyDigitsErr")
}

func BenchmarkAppendDigits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AppendDigits(text)
	}
}

func BenchmarkCopyDigits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CopyDigits(text)
	}
}

/*
goos: darwin
goarch: amd64
pkg: gocase/slice
BenchmarkAppendDigits
BenchmarkAppendDigits-4   	 1620531	       741 ns/op
BenchmarkCopyDigits
BenchmarkCopyDigits-4     	 1634118	       731 ns/op
PASS
*/
