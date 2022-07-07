/* higher order functions
assign functions as values to variables
*/

package main

import "fmt"

func main() {
	/*
		fn := func() {
			fmt.Println("fn invoked")
		}
	*/
	/*
		var fn = func() {
			fmt.Println("fn invoked")
		}
	*/

	var fn func()
	fn = func() {
		fmt.Println("fn invoked")
	}
	fn()

	var greet func(string)
	greet = func(userName string) {
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	}
	greet("Magesh")

	var add func(int, int) int
	add = func(x, y int) int {
		return x + y
	}
	fmt.Println(add(100, 200))
}
