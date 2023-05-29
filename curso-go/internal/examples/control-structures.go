package examples

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func TryControlStructures() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter a number: ")
	var s string
	if s, _ = reader.ReadString('\n'); true {
		s = strings.Trim(s, " \n")
	}
	number, _ := strconv.Atoi(s)

	// If
	if number > 10 {
		fmt.Printf("The number was greater than 10\n")
	} else {
		fmt.Printf("The number was lesser than 10\n")
	}

	// For basic loop
	for i := 0; i < number; i++ {
		fmt.Printf("Counting to %v, at %v\n", number, i+1)
	}

	// Switch
	switch number % 2 {
	case 0:
		fmt.Printf("Number %v is even\n", number)
	case 1:
		fmt.Printf("Number %v is odd\n", number)
	}

	// Switch without condition
	switch {
	case number < 10:
		fmt.Printf("Number %v is less than 10\n", number)
	case number < 20:
		fmt.Printf("Number %v is between 10 and 20\n", number)
	default:
		fmt.Printf("Number %v is greater than 20\n", number)
	}

	// For loop without all members
	j := 0
	for j < number {
		fmt.Printf("Counting to %v, at %v\n", number, j+1)
		j = j + 1
	}
	for {
		fmt.Printf("This message will loop until a break\n")
		break
	}
	for n := 0; n < number; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Printf("%v is odd\n", n)
	}

	// Declaration in condition
	if remainder := number % 5; remainder == 0 {
		fmt.Printf("%v is divisible by 5\n", number)
	} else {
		fmt.Printf("%v is not divisible by 5, with remainder of %v\n", number, remainder)
	}
}
