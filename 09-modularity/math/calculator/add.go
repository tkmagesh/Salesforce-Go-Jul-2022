package calculator

import "fmt"

func init() {
	fmt.Println("calculator initialized - add")
}

func Add(x, y int) int {
	operationCount++
	return x + y
}
