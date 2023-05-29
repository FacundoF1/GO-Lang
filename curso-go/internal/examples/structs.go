package examples

import "fmt"

func TryStructs() {

	// A simple struct
	type User struct {
		Name string
		Age  int
	}

	user := User{Name: "Sebatian", Age: 15}
	user.Age = 35
	fmt.Printf("The user is %v\n", user)
	fmt.Printf("The user %s is aged %v", user.Name, user.Age)
}
