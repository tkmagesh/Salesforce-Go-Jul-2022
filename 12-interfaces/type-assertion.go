package main

import "fmt"

func main() {
	//var x interface{}
	var x any

	//x = 10
	//x = "this is a string"
	//x = true
	//x = struct{}{}
	x = 10.98
	//y := x + 100

	if val, ok := x.(int); ok {
		y := val + 100
		fmt.Println(y)
	} else {
		fmt.Println("x is not an int")
	}

	switch val := x.(type) {
	case int:
		fmt.Println("x is an int, x + 100 = ", val+100)
	case string:
		fmt.Println("x is a string, len(x) = ", len(val))
	case bool:
		fmt.Println("x is a bool, !x =", !val)
	case struct{}:
		fmt.Println("x is an empty struct")
	default:
		fmt.Println("unknown type")
	}
}
