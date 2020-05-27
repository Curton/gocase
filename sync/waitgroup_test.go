package sync

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		// start new goroutine
		go func(i int) {
			defer wg.Done()
			time.Sleep(time.Millisecond * time.Duration(i))
			fmt.Printf("goroutine: %v end.\r\n", i)
		}(i)
	}

	// wait all goroutine done
	wg.Wait()
	fmt.Println("All goroutine done.")
}
