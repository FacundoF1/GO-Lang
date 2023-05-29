package examples

import "fmt"

func TryMaps() {

	m := make(map[int]string)
	m[200] = "OK"
	m[400] = "BAD_REQUEST"
	fmt.Printf("A map with values: %v\n", m)

	initializedMap := map[int]string{0: "Monday", 1: "Tuesday", 2: "Wednesday"}
	fmt.Printf("An initialized map: %v\n", initializedMap)

	for key, value := range m {
		fmt.Printf("Ranging over map, at key %v with value %v\n", key, value)
	}
}
