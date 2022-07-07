package calculator

import "fmt"

func init() {
	fmt.Println("calculator initialized - subtract")
}

func Subtract(x, y int) int {
	operationCount++
	return x - y
}
