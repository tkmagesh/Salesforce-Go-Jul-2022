package main

import "fmt"

type Employee struct {
	Id   int
	Name string
	City string
}

type FulltimeEmployee struct {
	Employee
	Name     string
	Benefits string
}

type Dummy struct {
	Name string
}

func main() {
	fte := FulltimeEmployee{
		Employee: Employee{
			Id:   100,
			Name: "Magesh",
			City: "Bengaluru",
		},
		Benefits: "Healthcare",
		Name:     "Suresh",
	}
	fmt.Println(fte.Employee.Id, fte.Employee.Name, fte.Employee.City, fte.Benefits)
	fmt.Println(fte.Id, fte.Name, fte.City, fte.Benefits)

}
