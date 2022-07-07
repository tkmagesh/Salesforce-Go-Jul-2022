/*
	Higher Order Functions
	- functions as return values
*/

package main

import (
	"fmt"
	"log"
	"time"
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

	/* logAdd(100, 200)
	logSubtract(100, 200) */

	/*
		profileAdd := getProfileOperation(add)
		profileAdd(100, 200)
	*/

	profileAndLogAdd := getProfileOperation(logAdd)
	profileAndLogSubtract := getProfileOperation(logSubtract)

	profileAndLogAdd(100, 200)
	profileAndLogSubtract(100, 200)
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

func getProfileOperation(operation func(int, int)) func(int, int) {
	return func(x, y int) {
		start := time.Now()
		operation(x, y)
		elpased := time.Now().Sub(start)
		fmt.Println("Elapsed : ", elpased)
	}
}

/* 3rd party library - can't modify them */
func add(x, y int) {
	time.Sleep(3 * time.Second)
	var result = x + y
	fmt.Println("Add result = ", result)
}

func subtract(x, y int) {
	time.Sleep(5 * time.Second)
	var result = x - y
	fmt.Println("Subtract result = ", result)
}
