package main

import "fmt"

func main() {
	//products := []string{"pen", "pencil", "marker"}

	//pre allocating the memory
	/*
		products := make([]string, 0, 20)
		products = append(products, "pen", "pencil", "marker")
	*/

	//pre allocating the memory & initializing
	products := make([]string, 3, 20)
	products[0] = "pen"
	products[1] = "pencil"
	products[2] = "marker"

	fmt.Printf("len = %d, capacity = %d, products = %v\n", len(products), cap(products), products)

	fmt.Println("Adding a new product")
	products = append(products, "notebook")
	fmt.Printf("len = %d, capacity = %d, products = %v\n", len(products), cap(products), products)

	fmt.Println("Adding a few more products")
	products = append(products, "book", "stapler", "scale")
	fmt.Printf("len = %d, capacity = %d, products = %v\n", len(products), cap(products), products)

	//creating a new slice pointing to the same array of products
	newProducts := products[0:2]

	//creating a new slice with its own array
	/*
		newProducts := make([]string, 2, 2)
		copy(newProducts, products)
	*/
	fmt.Println("newProducts =", newProducts)
	fmt.Println("After updating newProducts")
	newProducts[0] = "Fountain-Pen"
	fmt.Println("products =", products)
	fmt.Println("newProducts =", newProducts)
	fmt.Printf("len = %d, capacity = %d, newProducts = %v\n", len(newProducts), cap(newProducts), newProducts)

}
