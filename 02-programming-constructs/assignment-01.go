/*
Write a program that will print the prime numbers between start & end
accept the start & end from the user
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
