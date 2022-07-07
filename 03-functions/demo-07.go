/*
	Higher Order Functions
	- functions as return values
*/

package main

import (
	"fmt"
	"log"
)

func main() {
	fn := getFn()
	fn()

	/*
		add(100, 200)
		subtract(100, 200)
	*/

	/*
		logOperation(100, 200, add)
		logOperation(100, 200, subtract)
	*/
	logAdd := getLogOperation(add)
	logSubtract := getLogOperation(subtract)

	logAdd(100, 200)
	logSubtract(100, 200)

	getLogOperation(add)(1000, 2000)
}

func getFn() func() {
	return func() {
		fmt.Println("fn invoked")
	}
}

func getLogOperation(operation func(int, int)) func(int, int) {
	return func(x, y int) {
		log.Println("Before invocation")
		operation(x, y)
		log.Println("After invocation")
	}
}

/* 3rd party library - can't modify them */
func add(x, y int) {
	var result = x + y
	fmt.Println("Add result = ", result)
}

func subtract(x, y int) {
	var result = x - y
	fmt.Println("Subtract result = ", result)
}
