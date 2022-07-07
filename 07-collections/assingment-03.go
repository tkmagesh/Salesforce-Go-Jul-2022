/*
	Refactor the solution for assingment-02 using functions
*/
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

import (
	"errors"
	"fmt"
)

var operations map[int]func(int, int) int
var InvalidUserChoiceError = errors.New("Invalid user choice")

func init() {
	operations = map[int]func(int, int) int{
		1: add,
		2: subtract,
		3: multiply,
		4: divide,
	}
}

func main() {
	var userChoice int
	for {
		userChoice = getUserChoice()
		if userChoice == 5 {
			break
		}
		if err := process(userChoice); err == InvalidUserChoiceError {
			fmt.Println("Invalid choice")
		}
	}
	fmt.Println("Thank you!")
}

func getUserChoice() int {
	var userChoice int
	fmt.Println("Menu")
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")
	fmt.Println("5. Exit")

	fmt.Println("Enter your choice:")
	fmt.Scanln(&userChoice)
	return userChoice
}

func getOperands() (x, y int) {
	fmt.Println("Enter the values:")
	fmt.Scanln(&x, &y)
	return
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func multiply(x, y int) int {
	return x * y
}

func divide(x, y int) int {
	return x / y
}

func process(userChoice int) error {

	if operation, exists := operations[userChoice]; exists {
		x, y := getOperands()
		result := operation(x, y)
		fmt.Println("Result =", result)
		return nil
	}
	return InvalidUserChoiceError
}
