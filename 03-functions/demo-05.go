/*
Higher order functions
	- functions as arguments
*/

package main

import "fmt"

func main() {
	exec(fn)
}

func exec(f func()) {
	f()
}

func fn() {
	fmt.Println("fn invoked")
}
