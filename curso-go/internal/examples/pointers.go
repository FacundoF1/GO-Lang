package examples

import "fmt"

func TryPointers() {

	var age int = 34
	var agePointer *int = &age
	fmt.Println(age)
	*agePointer = 35
	fmt.Println(age)

}
