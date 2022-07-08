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
	for i := 0; i < 10; i++ {
		time.Sleep(500 * time.Millisecond)
		ch <- i * 10
	}
	close(ch)
}
