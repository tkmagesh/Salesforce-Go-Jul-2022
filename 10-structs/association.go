package main

import "fmt"

type Employee struct {
	Id   int
	Name string
	City string
	Org  *Organization
}

type Organization struct {
	Name    string
	Address string
}

func NewEmployee(id int, name string, city string, orgName string, orgAddress string) *Employee {
	return &Employee{
		Id:   id,
		Name: name,
		City: city,
		Org: &Organization{
			Name:    orgName,
			Address: orgAddress,
		},
	}
}

func main() {
	/*
		emp := Employee{
			Id:   100,
			Name: "Magesh",
			City: "Bengaluru",
			Org: Organization{
				Name:    "Salesforce",
				Address: "Bengaluru",
			},
		}
		fmt.Println(emp)
		fmt.Println(emp.Name)
		fmt.Println(emp.Org.Name)
	*/
	org := &Organization{
		Name:    "Salesforce",
		Address: "Bengaluru",
	}

	emp1 := Employee{
		Id:   100,
		Name: "Magesh",
		City: "Bengaluru",
		Org:  org,
	}

	emp2 := Employee{
		Id:   200,
		Name: "Suresh",
		City: "Bengaluru",
		Org:  org,
	}

	fmt.Printf("emp1 = %#v\n", emp1)
	fmt.Printf("emp2 = %#v\n", emp2)

	fmt.Println("After modifying the Org address")
	org.Address = "Newyork"
	fmt.Printf("org = %#v\n", org)
	fmt.Printf("emp1 org = %#v\n", emp1.Org)
	fmt.Printf("emp2 org = %#v\n", emp2.Org)

	emp3 := NewEmployee(300, "Ganesh", "Pune", "Salesforce", "Pune")
	fmt.Println(emp3)
}
