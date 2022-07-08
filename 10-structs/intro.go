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

	var emp = Employee{
		Id:   100,
		Name: "Magesh",
	}
	fmt.Println("emp Name = ", emp.Name)
	/*
		emp := Employee{
			Id:   100,
			Name: "Magesh",
		}
	*/

	//pointer
	//empPtr := new(Employee)
	empPtr := &Employee{
		Id:   100,
		Name: "Magesh",
	}
	fmt.Printf("%#v\n", empPtr)
	fmt.Println("emp Name = ", empPtr.Name)
}
