package main

import "fmt"

func main() {
	var no int = 100

	var noPtr *int = &no
	fmt.Println(noPtr)

	//dereferencing
	y := *noPtr
	fmt.Println(y)

	//no == *(&no)
	fmt.Println("Before incrementing, no = ", no)
	increment(&no)
	fmt.Println("After incrementing, no = ", no)

	x, y := 100, 200
	fmt.Println("Before swapping, ", x, y)
	swap(&x, &y)
	fmt.Println("After swapping, ", x, y)
}

func increment(no *int) {
	*no++
}

func swap(n1, n2 *int) {
	*n1, *n2 = *n2, *n1
}
