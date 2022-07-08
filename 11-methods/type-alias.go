package main

import "fmt"

//alias for string
type MyStr string

func (s MyStr) Lenght() int {
	return len(s)
}

func main() {

	//var s = MyStr("This is a string")
	var s = "This is a string"
	var myStr = MyStr(s)
	fmt.Println(myStr.Lenght())
}
