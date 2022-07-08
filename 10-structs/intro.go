package main

import "fmt"

type Employee struct {
	Id   int
	Name string
	City string
}

func main() {
	//var emp Employee
	//fmt.Println(emp)

	//var emp Employee = Employee{100, "Magesh", "Bangalore"}
	//var emp Employee = Employee{Id: 100, Name: "Magesh"}
	/*
		var emp Employee = Employee{
			Id:   100,
			Name: "Magesh",
		}
	*/
	/*
		var emp = Employee{
			Id:   100,
			Name: "Magesh",
		}
	*/
	/*
		emp := Employee{
			Id:   100,
			Name: "Magesh",
		}
	*/

	//pointer
	//emp := new(Employee)
	emp := &Employee{
		Id:   100,
		Name: "Magesh",
	}
	fmt.Printf("%#v\n", emp)
}
