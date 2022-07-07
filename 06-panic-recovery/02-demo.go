/* programmatically raising a panic */

package main

import (
	"errors"
	"fmt"
)

var DivideByZeroError = errors.New("divisor cannot be zero")

func main() {
	defer func() {
		fmt.Println("	deferred @ main")
		if err := recover(); err != nil {
			if err.(error) == DivideByZeroError {
				fmt.Println("dont try to divide by 0")
				return
			}
			fmt.Println("something went wrong!", err)
			return
		}
		fmt.Println("application shutdown successfully")
	}()
	q, r := divide(100, 0)
	fmt.Println(q, r)
}

func divide(x, y int) (quotient, remainder int) {
	if y == 0 {
		panic(DivideByZeroError)
	}
	quotient, remainder = x/y, x%y
	return
}
