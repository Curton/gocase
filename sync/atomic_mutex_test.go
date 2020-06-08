package sync

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"sync/atomic"
	"testing"
)

func BenchmarkAtomic(b *testing.B) {
	n := new(int32)
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(n, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	assert.Equal(b, int32(b.N), *n, "Err in BenchmarkAtomic")
}

func BenchmarkMutex(b *testing.B) {
	var n int32 = 0
	mux := sync.Mutex{}
	var wg sync.WaitGroup
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		go func() {
			mux.Lock()
			n++
			mux.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	assert.Equal(b, int32(b.N), n, "Err in BenchmarkMutex")
}

func BenchmarkMutex2(b *testing.B) {
	var n int32 = 0
	mux := sync.Mutex{}
	var wg sync.WaitGroup
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		go func() {
			mux.Lock()
			defer mux.Unlock()
			n++
			wg.Done()
		}()
	}
	wg.Wait()
	assert.Equal(b, int32(b.N), n, "Err in BenchmarkMutex2")
}

/*
goos: darwin
goarch: amd64
pkg: gocase/sync
BenchmarkAtomic
BenchmarkAtomic-16    	 2586414	       461 ns/op
BenchmarkMutex
BenchmarkMutex-16     	 2924312	       424 ns/op
BenchmarkMutex2
BenchmarkMutex2-16    	 2966654	       455 ns/op
PASS
*/
