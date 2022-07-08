package main

import (
	"fmt"
	"time"
)

/*

func main() {
	wg := &sync.WaitGroup{}
	var ch chan int = make(chan int)
	wg.Add(1)
	go add(100, 200, wg, ch)
	result := <-ch
	wg.Wait()
	fmt.Println(result)
}

func add(x, y int, wg *sync.WaitGroup, ch chan int) {
	time.Sleep(3 * time.Second)
	result := x + y
	ch <- result
	wg.Done()
}
*/

func main() {

	var ch chan int = make(chan int)
	go add(100, 200, ch)
	result := <-ch
	fmt.Println(result)
}

func add(x, y int, ch chan int) {
	time.Sleep(3 * time.Second)
	result := x + y
	ch <- result
}
