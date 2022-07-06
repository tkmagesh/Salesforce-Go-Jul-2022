/*
	Display the following menu
		1. Add
		2. Subtract
		3. Multiply
		4. Divide
		5. Exit

	Accept the user choice
	If the user choice is 1 - 4
		accept 2 numbers from the user
		perform the operation
		print the result
		display the menu again

	If the user choice = 5
		print "Thank you"
		exit the application

	If the user choice is NOT 1 - 5
		print "Invalid choice"
		display the menu again
*/

package main

import "fmt"

func main() {
	var x, y, userChoice, result int
	for {
		fmt.Println("Menu")
		fmt.Println("1. Add")
		fmt.Println("2. Subtract")
		fmt.Println("3. Multiply")
		fmt.Println("4. Divide")
		fmt.Println("5. Exit")

		fmt.Println("Enter your choice:")
		fmt.Scanln(&userChoice)
		if userChoice < 1 || userChoice > 5 {
			fmt.Println("Invalid choice...")
			continue
		}
		if userChoice == 5 {
			break
		}
		fmt.Println("Enter the values:")
		fmt.Scanln(&x, &y)
		switch userChoice {
		case 1:
			result = x + y
		case 2:
			result = x - y
		case 3:
			result = x * y
		case 4:
			result = x / y
		}
		fmt.Println("Result =", result)
	}
	fmt.Println("Thank you!")
}
