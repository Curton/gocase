package gopool

import (
	"fmt"
	"sync"
	"time"
)

func DemoFunc() {
	// do sth
	time.Sleep(time.Millisecond * time.Duration(10))
	fmt.Println("hello")
}

func SimplestPool() {
	start := time.Now()
	wg := new(sync.WaitGroup)
	dataCh := make(chan int, 100)

	for goId := 0; goId < 10; goId++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			// accept dataCh from dataCh
			for v := range dataCh {
				fmt.Printf("goroutine[%v]: %v\r\n", n, v)
			}
		}(goId)
	}

	for i := 0; i < 1000; i++ {
		dataCh <- i
	}
	close(dataCh)
	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}
