/* anonymous functions */

package main

import "fmt"

func main() {

	/* simple anonymous function */
	func() {
		fmt.Println("fn invoked")
	}()

	/* with arguments */
	func(userName string) {
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	}("Magesh")

	/* with return value */
	result := func(x, y int) int {
		return x + y
	}(100, 200)
	fmt.Println(result)

	/* with multiple return values */
	quotient, remainder := func(x, y int) (quotient, remainder int) {
		quotient, remainder = x/y, x%y
		return
	}(100, 7)
	fmt.Printf("quotient = %d, remainder = %d\n", quotient, remainder)
}
