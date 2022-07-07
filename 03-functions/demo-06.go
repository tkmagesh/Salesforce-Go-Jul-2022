package main

import (
	"fmt"
	"log"
)

func main() {
	/*
		logAdd(100, 200)
		logSubtract(100, 200)
	*/
	logOperation(100, 200, add)
	logOperation(100, 200, subtract)
}

func logOperation(x, y int, operation func(int, int)) {
	log.Println("Before invocation")
	operation(x, y)
	log.Println("After invocation")
}

/*
func logAdd(x, y int) {
	log.Println("Before invocation")
	add(x, y)
	log.Println("After invocation")
}

func logSubtract(x, y int) {
	log.Println("Before invocation")
	subtract(x, y)
	log.Println("After invocation")
}
*/

/* 3rd party library - can't modify them */
func add(x, y int) {
	var result = x + y
	fmt.Println("Add result = ", result)
}

func subtract(x, y int) {
	var result = x - y
	fmt.Println("Subtract result = ", result)
}
