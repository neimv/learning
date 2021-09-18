package main

import "fmt"

type Person struct {
	age int
}

type Employee struct {
	id       int
	name     string
	vacation bool
}

type FullTimeEmployee struct {
	Person
	Employee
	endDate string
}

type TemporaryEmployee struct {
	Person
	Employee
	taxRate int
}

type PrintInfo interface {
	GetMessage() string
}

func NewEmployee(id int, name string, vacation bool) *Employee {
	return &Employee{
		id:       id,
		name:     name,
		vacation: vacation,
	}
}

func (e *Employee) SetId(id int) {
	e.id = id
}

func (e *Employee) GetId() int {
	return e.id
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e *Employee) GetName() string {
	return e.name
}

func GetMessage(p PrintInfo) {
	fmt.Println(p.GetMessage())
}

func (ftEmployee FullTimeEmployee) GetMessage() string {
	return "Full Time Employee"
}

func (ftTemporary TemporaryEmployee) GetMessage() string {
	return "Temporary Employee"
}

func main() {
	e := Employee{}

	fmt.Printf("%v\n", e)

	e.id = 1
	e.name = "neimv"
	fmt.Printf("%v\n", e)
	e.SetId(5)
	e.SetName("Zephon")
	fmt.Printf("%v\n", e)

	fmt.Println(e.GetId())
	fmt.Println(e.GetName())

	e2 := Employee{
		id:       1,
		name:     "ax",
		vacation: true,
	}

	fmt.Printf("%v\n", e2)

	e3 := new(Employee)
	fmt.Printf("%v\n", *e3)
	e3.id = 2
	e3.name = "Zatara"
	fmt.Printf("%v\n", *e3)

	e4 := NewEmployee(1, "hello", false)
	fmt.Printf("%v\n", *e4)

	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "name"
	ftEmployee.age = 2
	ftEmployee.id = 5
	tEmployee := TemporaryEmployee{}

	fmt.Printf("%v\n", ftEmployee)
	GetMessage(ftEmployee)
	GetMessage(tEmployee)
}
