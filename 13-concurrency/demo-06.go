package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	if count, err := strconv.Atoi(os.Args[1]); err == nil {
		wg := &sync.WaitGroup{}
		for i := 1; i <= count; i++ {
			wg.Add(1)
			go f1(i, wg)
		}
		f2()
		wg.Wait()
	} else {
		fmt.Println("Usage : demo-06 <count>")
	}
	fmt.Scanln()
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
