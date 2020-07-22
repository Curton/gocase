/**
 * @author Covey Liu, covey@liukedun.com
 * @date 2020/7/22 11:14
 */

package _chan

import (
	"fmt"
	"testing"
	"time"
)

func Test(t *testing.T) {
	ch := make(chan bool)
	ch2 := make(chan bool)

	go func() {
		defer fmt.Println("go routine exit")
		for {
			select {
			case <-ch:
				fmt.Println("received from ch")
			case b := <-ch2:
				fmt.Println("received from ch2 ", b)
				return
			}
		}
	}()

	ch <- true

	close(ch2) // will send inf false to chan, A closed channel is always available for read

	time.Sleep(time.Hour)
}
