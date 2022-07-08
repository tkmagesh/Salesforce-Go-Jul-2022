package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go fn(ch)
	fmt.Println("attempting to receive 10")
	fmt.Println(<-ch)
	fmt.Println("Successfully received 10")

	fmt.Println("attempting to receive 20")
	fmt.Println(<-ch)
	fmt.Println("Successfully received 20")

	fmt.Println("attempting to receive 30")
	fmt.Println(<-ch)
	fmt.Println("Successfully received 30")

	fmt.Println("attempting to receive 40")
	fmt.Println(<-ch)
	fmt.Println("Successfully received 40")

	fmt.Println("attempting to receive 50")
	fmt.Println(<-ch)
	fmt.Println("Successfully received 50")

}

func fn(ch chan int) {
	fmt.Println("	Attempting to send 10")
	ch <- 10
	fmt.Println("	Successfully sent 10")
	time.Sleep(500 * time.Millisecond)

	fmt.Println("	Attempting to send 20")
	ch <- 20
	fmt.Println("	Successfully sent 20")
	time.Sleep(500 * time.Millisecond)

	fmt.Println("	Attempting to send 30")
	ch <- 30
	fmt.Println("	Successfully sent 30")
	time.Sleep(500 * time.Millisecond)

	fmt.Println("	Attempting to send 40")
	ch <- 40
	fmt.Println("	Successfully sent 40")
	time.Sleep(500 * time.Millisecond)

	fmt.Println("	Attempting to send 50")
	ch <- 50
	fmt.Println("	Successfully sent 50")
	time.Sleep(500 * time.Millisecond)

}
