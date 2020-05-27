package gopool

import (
	"strconv"
	"sync"
)

func DemoFunc(v int) {
	// do sth
	strconv.Itoa(v)
}

func SimplestPool() {
	wg := new(sync.WaitGroup)
	dataCh := make(chan int, 100)

	for goId := 0; goId < 10; goId++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			// accept dataCh from dataCh
			for v := range dataCh {
				// time.Sleep(time.Microsecond * 100)
				// fmt.Printf("goroutine[%v]: %v\r\n", n, v)
				DemoFunc(v)
			}
		}(goId)
	}

	for i := 0; i < 100000; i++ {
		dataCh <- i
	}
	close(dataCh)
	wg.Wait()
}

func NoPool() {
	wg := new(sync.WaitGroup)

	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			// time.Sleep(time.Microsecond * 100)
			//fmt.Println("goroutine", n)
			DemoFunc(n)
		}(i)
	}
	wg.Wait()
}
