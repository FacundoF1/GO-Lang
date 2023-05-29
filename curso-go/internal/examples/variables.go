package examples

import "fmt"

func TryVariables() {

	// Defining a variable
	var count int

	// Defining a constant
	const max int = 10

	// Defining multiple variables
	var (
		name    string
		amount  float32
		enabled bool
	)

	// Defining with type inference
	const cuil = "20334445550"
	nationality := "arg"

	fmt.Printf("count=%v, max=%v, name=%v, amount=%v, enabled=%v, cuil=%v, nationality=%v\n", count, max, name, amount, enabled, cuil, nationality)
}
