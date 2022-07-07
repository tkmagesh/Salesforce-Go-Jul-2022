package calculator

import "fmt"

var operationCount int

func OpCount() int {
	return operationCount
}

func init() {
	fmt.Println("calculator initialized - 1")
}

func init() {
	fmt.Println("calculator initialized - 2")
}
