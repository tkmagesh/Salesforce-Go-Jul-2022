package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go fn(ch)
	for no := range ch {
		fmt.Println(no)
	}
	fmt.Println("Channel is closed")
}

func fn(ch chan int) {
	i := 0
	timeoutCh := time.After(10 * time.Second)
LOOP:
	for {
		select {
		case ch <- i * 10:
			time.Sleep(500 * time.Millisecond)
			i++
		case <-timeoutCh:
			fmt.Println("timeout occured")
			break LOOP
		}
	}
	close(ch)
}

/* func timeout(d time.Duration) <-chan time.Time {
	timeoutCh := make(chan time.Time)
	go func() {
		time.Sleep(d)
		timeoutCh <- time.Now()
	}()
	return timeoutCh
} */
