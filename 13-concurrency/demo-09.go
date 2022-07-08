package main

import (
	"fmt"
	"sync"
)

var counter int
var mutex sync.Mutex

func main() {
	wg := &sync.WaitGroup{}
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go f1(i, wg)
	}
	wg.Wait()
	fmt.Println("Done")
	fmt.Println("Counter = ", counter)
}

func f1(id int, wg *sync.WaitGroup) {
	defer wg.Done() //decrements the counter by 1
	//fmt.Printf("f1-[%d] started\n", id)
	//time.Sleep(time.Duration(rand.Intn(5000)) * time.Millisecond)
	mutex.Lock()
	counter++
	mutex.Unlock()
	//fmt.Printf("f1-[%d] completed\n", id)
}
