package main

import (
	"fmt"
	"reflect"
)

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

	//Complex
	var c complex64 = 10 + 2i
	fmt.Printf("%T\n", c)
	fmt.Println("real = ", real(c))
	fmt.Println("imaginary = ", imag(c))

	//Rune
	rune1 := 'B'
	rune2 := 'g'
	rune3 := '\a'

	// Displaying rune and its type
	fmt.Printf("Rune 1: %c; Unicode: %U; Type: %s\n", rune1,
		rune1, reflect.TypeOf(rune1))

	fmt.Printf("Rune 2: %c; Unicode: %U; Type: %s\n", rune2,
		rune2, reflect.TypeOf(rune2))

	fmt.Printf("Rune 3: Unicode: %U; Type: %s\n", rune3,
		reflect.TypeOf(rune3))

	//type conversion
	var i int = 100
	var f float32 = float32(i)
	fmt.Println("f = ", f)

	var f2 float32 = 1000.999
	var no int = int(f2)
	fmt.Println(no)

	//For string to number conversion and vice versa , use strconv package
	/*
		x1 := "100"
		y1 := strconv.Atoi(x1) + 100 //TO BE FIXEX
	*/

	//Constant
	//const pi float32 = 3.14
	const pi = 31.4
	//pi = 2
	fmt.Println(pi)

	//iota
	/*
		const (
			red   = iota
			green = iota
			blue  = iota
		)
	*/

	/*
		const (
			red = iota
			green
			blue
		)
	*/

	/*
		const (
			red = iota + 2
			green
			blue
		)
	*/

	/*
		const (
			red = iota * 2
			green
			blue
		)
	*/

	const (
		red = iota * 2
		green
		_
		blue
	)
	fmt.Println("Red =", red, "Green =", green, "Blue =", blue)

	const (
		VERBOSE = 1 << iota
		CONFIG_FROM_DISK
		DATABASE_REQUIRED
		LOGGER_ACTIVATED
		DEBUG
		FLOAT_SUPPORT
		RECOVERY_MODE
		REBOOT_ON_FAILURE
	)
	fmt.Printf("%b, %b, %b, %b, %b, %b, %b, %b\n", VERBOSE, CONFIG_FROM_DISK, DATABASE_REQUIRED, LOGGER_ACTIVATED, DEBUG, FLOAT_SUPPORT, RECOVERY_MODE, REBOOT_ON_FAILURE)

}
