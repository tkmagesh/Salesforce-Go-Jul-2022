package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	stop := make(chan bool)
	go fn(ch, stop)

	go func() {
		fmt.Println("Hit ENTER to stop")
		fmt.Scanln()
		stop <- true
	}()

	for no := range ch {
		fmt.Println(no)
	}

	fmt.Println("Channel is closed")
}

func fn(ch chan int, stop <-chan bool) {
	i := 0

LOOP:
	for {
		select {
		case ch <- i * 10:
			time.Sleep(500 * time.Millisecond)
			i++
		case <-stop:
			fmt.Println("stop requested")
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
