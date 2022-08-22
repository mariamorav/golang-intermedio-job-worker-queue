package main

import "fmt"

type Employee struct {
	id   int
	name string
}

func (e *Employee) SetId(id int) {
	e.id = id
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e *Employee) GetId() int {
	return e.id
}

func (e *Employee) GetName() string {
	return e.name
}

func main() {
	employee := Employee{}
	fmt.Printf("%v\n", employee)
	employee.id = 0
	employee.name = "Maria"
	fmt.Printf("%v\n", employee)
	employee.SetId(5)
	employee.SetName("Ana")
	fmt.Println(employee.GetId())
	fmt.Println(employee.GetName())

}
