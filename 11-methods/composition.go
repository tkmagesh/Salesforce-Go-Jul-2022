package main

import "fmt"

type Employee struct {
	Id   int
	Name string
	City string
}

func (emp Employee) Format() string {
	return fmt.Sprintf("Id = %d, Name = %q, City = %q", emp.Id, emp.Name, emp.City)
}

type FulltimeEmployee struct {
	Employee
	Benefits string
}

//overriding the Format() method of Employee

func (fte FulltimeEmployee) Format() string {
	return fmt.Sprintf("%v, Benefits = %q", fte.Employee.Format(), fte.Benefits)
}

func main() {
	fte := FulltimeEmployee{
		Employee: Employee{
			Id:   100,
			Name: "Magesh",
			City: "Bengaluru",
		},
		Benefits: "Healthcare",
	}
	/*
		fmt.Println(fte.Employee.Id, fte.Employee.Name, fte.Employee.City, fte.Benefits)
		fmt.Println(fte.Id, fte.Name, fte.City, fte.Benefits)
	*/

	fmt.Println(fte.Format()) //Format method inherited from Employee
	fmt.Println(fte.Employee.Format())
}
