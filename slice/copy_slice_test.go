package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var text = []byte("ashfdhkjadfbkjbvjfdsbkgbkdfjdsfbhajskdfghkabhlfd")
var text2 = []byte("")

func TestAppend(t *testing.T) {
	assert.Equal(t, text, Append(text), "TestAppend Err 1")
	assert.Equal(t, text2, Append(text2), "TestAppend Err 2")
	assert.Equal(t, []byte(nil), Append(nil), "TestAppend Err 3")
}

func TestAppend2(t *testing.T) {
	assert.Equal(t, text, Append2(text), "TestAppend2 Err 1")
	assert.Equal(t, text2, Append2(text2), "TestAppend2 Err 2")
	assert.Equal(t, []byte(nil), Append2(nil), "TestAppend2 Err 3")
}

func TestAppend3(t *testing.T) {
	assert.Equal(t, text, Append3(text), "TestAppend3 Err 1")
	assert.Equal(t, text2, Append3(text2), "TestAppend3 Err 2")
	assert.Equal(t, []byte(nil), Append3(nil), "TestAppend3 Err 3")
}

func TestCopy(t *testing.T) {
	assert.Equal(t, text, Copy(text), "CopyDigits Err")
	assert.Equal(t, text2, Copy(text2), "CopyDigits Err")
	assert.Equal(t, []byte(nil), Copy(nil), "CopyDigits Err")
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

func BenchmarkAppendDigits3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Append3(text)
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
BenchmarkAppendDigits
BenchmarkAppendDigits-16     	34813472	        32.8 ns/op
BenchmarkAppendDigits2
BenchmarkAppendDigits2-16    	42832412	        28.2 ns/op
BenchmarkAppendDigits3
BenchmarkAppendDigits3-16    	37208007	        32.1 ns/op
BenchmarkCopyDigits
BenchmarkCopyDigits-16       	42077460	        28.5 ns/op
PASS
*/
