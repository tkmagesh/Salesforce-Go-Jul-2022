package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go f1(i, wg)
	}
	f2()
	wg.Wait()
}

func f1(id int, wg *sync.WaitGroup) {
	defer wg.Done() //decrements the counter by 1
	fmt.Printf("f1-[%d] started\n", id)
	time.Sleep(time.Duration(rand.Intn(5000)) * time.Millisecond)
	fmt.Printf("f1-[%d] completed\n", id)
}

func f2() {
	fmt.Println("f2 invoked")
}
