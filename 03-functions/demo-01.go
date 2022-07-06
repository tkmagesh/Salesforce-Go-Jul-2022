/* function basics */

package main

import "fmt"

func main() {
	fn()
	greet("Magesh")
	fmt.Print(getGreetMsg("Suresh"))
	fmt.Printf("Adding 100 & 200, result = %d\n", add(100, 200))
	//fmt.Println(divide(100, 7))
	/*
		q, r := divide(100, 7)
		fmt.Printf("Dividing 100 by 7, quotient = %d and remainder = %d\n", q, r)
	*/

	/*
		q, _ := divide(100, 7)
		fmt.Printf("Dividing 100 by 7, quotient = %d\n", q)
	*/

	_, r := divide(100, 7)
	fmt.Printf("Dividing 100 by 7, remainder = %d\n", r)
}

func fn() {
	fmt.Println("fn invoked")
}

/* 1 argument */
func greet(userName string) {
	fmt.Printf("Hi %s, Have a nice day!\n", userName)
}

/* 1 argument and 1 return value */
func getGreetMsg(userName string) string {
	return fmt.Sprintf("Hi %s, Have a nice day!\n", userName)
}

/* 2 arguments and 1 return value */
/*
func add(x int, y int) int {
	return x + y
}
*/

/*
func add(x, y int) int {
	return x + y
}
*/

//named return result
func add(x, y int) (result int) {
	result = x + y
	return
}

/* 2 arguments and 2 return values */
/*
func divide(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}
*/

/*
func divide(x, y int) (quotient int, remainder int) {
	quotient = x / y
	remainder = x % y
	return
}
*/

func divide(x, y int) (quotient, remainder int) {
	/*
		quotient = x / y
		remainder = x % y
	*/
	quotient, remainder = x/y, x%y
	return
}
