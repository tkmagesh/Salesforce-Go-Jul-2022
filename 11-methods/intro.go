package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

func (p Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %v, Units = %d, Category = %q", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

func (p *Product) ApplyDiscount(discountPercent float32) {
	p.Cost = p.Cost * ((100 - discountPercent) / 100)
}

func main() {
	pen := Product{
		Id:       100,
		Name:     "Pen",
		Cost:     10,
		Units:    100,
		Category: "Stationary",
	}
	//fmt.Println(Format(pen))
	fmt.Println(pen.Format())
	fmt.Println("After applying 10% discount")
	//ApplyDiscount(&pen, 10)

	//(&pen).ApplyDiscount(10)
	pen.ApplyDiscount(10)

	//fmt.Println(Format(pen))
	fmt.Println(pen.Format())
}
