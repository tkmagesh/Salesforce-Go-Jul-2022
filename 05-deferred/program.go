package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("	deferred - [main]")
	}()
	fmt.Println("[main] started")
	fmt.Println("f1() = ", f1())
	fmt.Println("[main] completed")
}

func f1() (result int) {
	var no int = 100
	defer func(no int) {
		fmt.Println("	deferred - [f1] - 1, no = ", no)
		result = 2000
	}(no)
	defer func() {
		fmt.Println("	deferred - [f1] - 2")
	}()
	defer func() {
		fmt.Println("	deferred - [f1] - 3")
	}()
	fmt.Println("[f1] started")
	f2()
	fmt.Println("[f1] completed")
	no = 200
	result = 200
	return
}

func f2() {
	defer func() {
		fmt.Println("	deferred - [f2]")
	}()
	fmt.Println("[f2] started")
	var divisor int = 0
	fmt.Println(100 / divisor)
	fmt.Println("[f2] completed")
}
