/* if, for, switch */
package main

import "fmt"

func main() {

	//if else
	/*
		no := 21
		if no%2 == 0 {
			fmt.Printf("%d is an even number\n", no)
		} else {
			fmt.Printf("%d is an odd number\n", no)
		}
	*/

	if no := 21; no%2 == 0 { //=> the scope of the 'no' variable is limited to the 'if else' block
		fmt.Printf("%d is an even number\n", no)
	} else {
		fmt.Printf("%d is an odd number\n", no)
	}

	//switch case
	/*
		score 0 - 3 => "Terrible"
		score 4 - 7 => "Not Bad"
		score 8 - 9 => "Very Good"
		score 10 => "Excellent"
		otherwise => "Invalid score"
	*/

	/*
		switch score := 7; score {
		case 0:
			fmt.Println("Terrible")
		case 1:
			fmt.Println("Terrible")
		case 2:
			fmt.Println("Terrible")
		case 3:
			fmt.Println("Terrible")
		case 4:
			fmt.Println("Not Bad")
		case 5:
			fmt.Println("Not Bad")
		case 6:
			fmt.Println("Not Bad")
		case 7:
			fmt.Println("Not Bad")
		case 8:
			fmt.Println("Very Good")
		case 9:
			fmt.Println("Very Good")
		case 10:
			fmt.Println("Excellent")
		default:
			fmt.Println("Invalid Score")

		}
	*/
	switch score := 7; score {
	case 0, 1, 2, 3:
		fmt.Println("Terrible")
	case 4, 5, 6, 7:
		fmt.Println("Not Bad")
	case 8, 9:
		fmt.Println("Very Good")
	case 10:
		fmt.Println("Excellent")
	default:
		fmt.Println("Invalid Score")
	}

	x := 11
	switch {
	case x%2 == 0:
		fmt.Printf("%d is even", x)
	case x%2 != 0:
		fmt.Printf("%d is odd", x)
	}

	//fallthrough
	fmt.Println("switch case with fallthrough")
	n := 4
	switch n {
	case 0:
		fmt.Println("is zero")
	case 1:
		fmt.Println("is <= 1")
		fallthrough
	case 2:
		fmt.Println("is <= 2")
		fallthrough
	case 3:
		fmt.Println("is <= 3")
		fallthrough
	case 4:
		fmt.Println("is <= 4")
		fallthrough
	case 5:
		fmt.Println("is <= 5")
		fallthrough
	case 6:
		fmt.Println("is <= 6")
		fallthrough
	case 7:
		fmt.Println("is <= 7")
		//fallthrough
	case 8:
		fmt.Println("is <= 8")
	}

	plan := "Pro"
	switch plan {
	case "Super":
		fmt.Println("All [Super] features included")
		fallthrough
	case "Premium":
		fmt.Println("All [Premium] features included")
		fallthrough
	case "Pro":
		fmt.Println("All [Pro] features included")
		fallthrough
	case "Free":
		fmt.Println("All [Free] features included")
	}

	//LOOPs using 'for'
	fmt.Println()
	fmt.Println("LOOPs using for")
	fmt.Println("for (v1)")
	for i := 1; i <= 5; i++ {
		fmt.Printf("i = %d\n", i)
	}

	fmt.Println("for (v2) [imitating while]")
	i := 1
	for i <= 5 {
		fmt.Printf("i = %d\n", i)
		i++
	}

	fmt.Println("for (v3) [infinite loop]")
	numSum := 1
	for {
		numSum += numSum
		fmt.Printf("numSum = %d\n", numSum)
		if numSum > 100 {
			break
		}
	}

	fmt.Println()
	fmt.Println("Using LABELs")
LOOP:
	for i := 1; i <= 10; i++ {
		for j := 1; j < 10; j++ {
			fmt.Printf("i = %d, j = %d\n", i, j)
			if i == j {
				continue LOOP
			}
		}
	}
}
