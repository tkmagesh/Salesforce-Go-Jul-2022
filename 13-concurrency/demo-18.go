package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)

	go f1(ch1)
	go f2(ch2)

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println(<-ch3)
	}()

	/*
		fmt.Println(<-ch1)
		fmt.Println(<-ch2)
	*/

	for i := 0; i < 3; i++ {
		select {
		case d1 := <-ch1:
			fmt.Println(d1)
		case d2 := <-ch2:
			fmt.Println(d2)
		case ch3 <- "data from ch3":
			fmt.Println("sent data to ch3")
		}
	}
}

func f1(ch chan string) {
	time.Sleep(4 * time.Second)
	ch <- "f1 done"
}

func f2(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "f2 done"
}
