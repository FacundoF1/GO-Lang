package examples

import "fmt"

// Function with simlpe signature
func BasicFunction(name string) string {
	return "hello " + name
}

// Function with multiple return values
func MultiReturnFunction(name string) (string, int, error) {
	return "Seba", 35, nil
}

// Function as variable
func Lambdas() {
	lambda := func() {
		fmt.Println("This is a lambda")
	}
	var sameLambda func() = lambda
	lambda()
	sameLambda()
}

// Function with variable parameters
func VariadicFunction(params ...string) {
	for _, param := range params {
		fmt.Println(param)
	}
}

func VariadicFunctionCaller() {
	VariadicFunction("param1", "param2", "param3")
}

type User struct {
	Name string
}

func (u User) Greet() {
	fmt.Printf("Hello, my name is %v\n", u.Name)
}

func (u *User) Rename(newName string) {
	u.Name = newName
}

func TryFunctions() {
	greeting := BasicFunction("seba")
	fmt.Println(greeting)
	name, age, err := MultiReturnFunction("Seba")
	fmt.Printf("Hello, my name is %s, my age is %v, and this is my error: %v\n", name, age, err)
	Lambdas()
	VariadicFunctionCaller()
	user := User{"Seba"}
	user.Greet()
	user.Rename("Sebastian")
	user.Greet()
}
