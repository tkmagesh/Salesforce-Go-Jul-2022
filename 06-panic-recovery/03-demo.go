package main

import "fmt"

func main() {
	q, r, err := divideClient(100, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(q, r)
}

func divideClient(x, y int) (quotient, remainder int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
			return
		}
	}()
	quotient, remainder = divide(x, y)
	return
}

//third party api
func divide(x, y int) (quotient, remainder int) {
	quotient, remainder = x/y, x%y
	return
}
