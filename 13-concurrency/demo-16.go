package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go fn(ch)
	for {
		if no, ok := <-ch; ok {
			fmt.Println(no)
		} else {
			fmt.Println("Channel is closed")
			break
		}
	}
}

func fn(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i * 10
	}
	close(ch)
}
