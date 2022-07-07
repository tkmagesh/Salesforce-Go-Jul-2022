package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("	deferred @ main")
		if err := recover(); err != nil {
			fmt.Println("something went wrong!", err)
			return
		}
		fmt.Println("application shutdown successfully")
	}()
	q, r := divide(100, 7)
	fmt.Println(q, r)
}

func divide(x, y int) (quotient, remainder int) {
	quotient, remainder = x/y, x%y
	return
}
