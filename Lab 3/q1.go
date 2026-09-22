package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Job    string
	Salary float64
}

func (p *Person) ReadData() {
	fmt.Print("Enter Name:")
	fmt.Scanln(&p.Name)

	fmt.Print("Enter Age:")
	fmt.Scanln(&p.Age)

	fmt.Print("Enter Job:")
	fmt.Scanln(&p.Job)

	fmt.Print("Enter Salary:")
	fmt.Scanln(&p.Salary)
}

func (p Person) PrintDetails() {
	fmt.Println("----------------------------")
	fmt.Printf("Name   : %s\n", p.Name)
	fmt.Printf("Age    : %d\n", p.Age)
	fmt.Printf("Job    : %s\n", p.Job)
	fmt.Printf("Salary : %.2f\n", p.Salary)
	fmt.Println("----------------------------")
}

func main() {
	var p1, p2 Person

	fmt.Println("Enter details for Person 1:")
	p1.ReadData()

	fmt.Println("\nEnter details for Person 2:")
	p2.ReadData()

	fmt.Println("\nPerson 1 Details:")
	p1.PrintDetails()

	fmt.Println("\nPerson 2 Details:")
	p2.PrintDetails()
}
