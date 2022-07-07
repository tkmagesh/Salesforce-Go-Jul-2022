package main

import "fmt"

func main() {
	//var nos []int
	/*
		var nos []int = make([]int, 5)
		nos[0] = 3
	*/
	var nos []int = []int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	fmt.Println()
	fmt.Println("Iterating using indexer")
	for idx := 0; idx < len(nos); idx++ {
		fmt.Printf("nos[%d] = %d\n", idx, nos[idx])
	}

	fmt.Println()
	fmt.Println("Iterating using range")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}

	fmt.Println()
	fmt.Println("Adding a new value")
	nos = append(nos, 10)
	fmt.Println(nos)

	fmt.Println("Adding more than one value")
	nos = append(nos, 20, 30, 40)
	fmt.Println(nos)

	fmt.Println("Adding values from another slice")
	hundreds := []int{100, 200, 300, 400}
	nos = append(nos, hundreds...)
	fmt.Println(nos)

	fmt.Println("slicing")
	/*
		[lo : hi] => from lo to hi-1
		[lo : ] => from lo to end
		[ : hi] => from 0 to hi-1
	*/
	fmt.Println("nos[2:6] => ", nos[2:6])
	fmt.Println("nos[2:] => ", nos[2:])
	fmt.Println("nos[:6] => ", nos[:6])

}
