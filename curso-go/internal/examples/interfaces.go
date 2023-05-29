package examples

import (
	"fmt"
)

type Greeter interface {
	Greet() string
}

type Person string

func (p Person) Greet() string {
	return fmt.Sprintf("Hello, my name is %s", p)
}

type Company struct {
	Name    string
	CeoName string
}

func (c Company) Greet() string {
	return fmt.Sprintf("Hello, welcome to %s, I'm the CEO %s", c.Name, c.CeoName)
}

func PrintGreet(greeter Greeter) {
	fmt.Println(greeter.Greet())
}

func PrintThing(thing interface{}) {
	fmt.Println(thing)
}

func PrintType(thing interface{}) {
	switch thing.(type) {
	case Person:
		fmt.Println("The thing is a person")
	case Company:
		fmt.Println("The thing is a company")
	}
}

func TryInterfaces() {

	var person Person = "Sebastian"
	var company Company = Company{Name: "MODO", CeoName: "Rafael Soto"}
	PrintGreet(person)
	PrintGreet(company)
	PrintThing(person)
	PrintThing(company)
	PrintType(person)
	PrintType(company)

}
