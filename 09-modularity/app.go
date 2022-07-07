package main

import (
	"fmt"
	//"modularity-demo/calculator"

	_ "modularity-demo/math/calculator"
	"modularity-demo/math/utils"

	"github.com/fatih/color"
)

func init() {
	fmt.Println("app initialized")
}
func main() {
	fmt.Println("app executed")
	/*
		fmt.Println(calculator.Add(100, 200))
		fmt.Println(calculator.Subtract(100, 200))
	*/
	/*
		fmt.Println(calc.Add(100, 200))
		fmt.Println(calc.Subtract(100, 200))
		fmt.Println("Op count = ", calc.OpCount())
	*/
	fmt.Println(utils.IsEven(20))

	color.Blue("This text will be printed in blue color")
	color.Red("This text will be printed in red color")

}
