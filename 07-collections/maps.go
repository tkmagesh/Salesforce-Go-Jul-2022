package main

import "fmt"

func main() {
	/*
		//var productRanks map[string]int = make(map[string]int)
		var productRanks map[string]int = map[string]int{}
		productRanks["pen"] = 5
	*/
	//var productRanks map[string]int = map[string]int{"pen": 5, "pencil": 1}
	var productRanks map[string]int = map[string]int{
		"pen":    5,
		"pencil": 1,
	}
	fmt.Println(productRanks)

	productRanks["marker"] = 4
	productRanks["scribble-pad"] = 2
	fmt.Println(productRanks)

	fmt.Println("Iterating a map")
	for product, rank := range productRanks {
		fmt.Printf("Rank of %q is %d\n", product, rank)
	}

	fmt.Println()
	fmt.Println("Deleting an item")
	delete(productRanks, "abc")
	fmt.Println(productRanks)

	fmt.Println()
	fmt.Println("Checking for the existing of a key")
	keyToCheck := "abc"
	if rank, exists := productRanks[keyToCheck]; exists {
		fmt.Printf("key(%q) exists with rank %d\n", keyToCheck, rank)
	} else {
		fmt.Printf("key(%q) does not exist\n", keyToCheck)
	}

}
