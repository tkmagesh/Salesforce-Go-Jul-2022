package main

import (
	"fmt"
	"time"
)

func main() {
	go f1() //scheduling the execution of the function
	f2()
	time.Sleep(1 * time.Millisecond)
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("f2 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
