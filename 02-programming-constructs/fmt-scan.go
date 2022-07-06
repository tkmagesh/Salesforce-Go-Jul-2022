package main

import "fmt"

func main() {
	/*
		var name string
		fmt.Println("Enter your name:")
		fmt.Scan(&name)
		fmt.Printf("Hi %s, Have a nice day!\n", name)
	*/

	/*
		var no int
		fmt.Println("Enter a number:")
		fmt.Scan(&no)
		fmt.Println("The given number is :", no)
	*/

	var firstName, lastName string
	fmt.Println("Enter your firstName & lastName:")
	fmt.Scanln(&firstName, &lastName)
	fmt.Println("firstName = ", firstName, "lastName =", lastName)

}
