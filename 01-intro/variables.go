package main

import "fmt"

//var x int = 100
var x = 100 //unused variables at the package level are allowed (not in a function)

func main() {
	/*
		var s string
		s = "Hello World!"
	*/

	/*
		var s string = "Hello World!"
	*/

	//type inference
	/*
		var s = "Hello World!"
	*/

	s := "Hello World!" // := syntax can be use ONLY in function scope
	fmt.Println(s)

	/* handling multiple variables */
	/*
		var x int
		var y int
		var str string
		var result int
		x = 100
		y = 200
		str = "Result = "
		result = x + y
		fmt.Println(str, result)
	*/

	/*
		var x, y, result int
		var str string
		x = 100
		y = 200
		str = "Result = "
		result = x + y
		fmt.Println(str, result)
	*/

	/*
		var (
			x, y, result int
			str          string
		)
		x = 100
		y = 200
		str = "Result = "
		result = x + y
		fmt.Println(str, result)
	*/

	/*
		var x int = 100
		var y int = 200
		var str string = "Result = "
		var result int = x + y
		fmt.Println(str, result)
	*/

	/*
		var x, y int = 100, 200
		var str string = "Result = "
		var result int = x + y
		fmt.Println(str, result)
	*/

	/*
		var (
			x, y   int    = 100, 200
			str    string = "Result = "
			result int    = x + y
		)
		fmt.Println(str, result)
	*/

	/*
		var (
			x, y   = 100, 200
			str    = "Result = "
			result = x + y
		)
		fmt.Println(str, result)
	*/

	/*
		var (
			x, y, str = 100, 200, "Result = "
			result    = x + y
		)
		fmt.Println(str, result)
	*/

	x, y, str := 100, 200, "Result = "
	result := x + y
	fmt.Println(str, result)

}
