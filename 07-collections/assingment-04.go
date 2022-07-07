/*
Write a program that will print the prime numbers between start & end
accept the start & end from the user

refactor the program to encapsulate the prime number generation logic in a function
the main function should be responsible for accpeting inputs and printing the result
*/

package main

import "fmt"

func main() {
	var start, end int
	fmt.Println("Enter the start and end:")
	fmt.Scanln(&start, &end)
LOOP:
	for no := start; no <= end; no++ {
		for i := 2; i <= (no / 2); i++ {
			if no%i == 0 {
				continue LOOP
			}
		}
		fmt.Println("Prime No : ", no)
	}
}
