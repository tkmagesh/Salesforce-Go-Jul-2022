/* variadic functions */

package main

import "fmt"

func main() {
	fmt.Println(sum(10))
	fmt.Println(sum(10, 20))
	fmt.Println(sum(10, 20, 30, 40, 50))
	fmt.Println(sum())
}

func sum(nos ...int) int {
	var result int
	for idx := 0; idx < len(nos); idx++ {
		result += nos[idx]
	}
	return result
}
