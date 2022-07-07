package main

import (
	"errors"
	"fmt"
)

var DivideByZeroError = errors.New("divisor cannot be zero")

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("something went wrong : ", err)
			return
		}
		fmt.Println("Thank you!")
	}()
	fmt.Println("main started")
	quotient, remainder, err := divideClient(100, 0)
	if err == DivideByZeroError {
		fmt.Println("Please dont attempt to divide by zero.")
		return
	}
	fmt.Println(quotient, remainder)
	fmt.Println("main completed")
}

func divideClient(x, y int) (quotient, remainder int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
			if err.Error() == "runtime error: integer divide by zero" {
				err = DivideByZeroError
			}
		}
	}()
	quotient, remainder = divide(x, y)
	return
}

//3rd party
func divide(x, y int) (quotient, remainder int) {
	quotient = x / y
	remainder = x % y
	return
}
