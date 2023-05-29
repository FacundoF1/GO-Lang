package examples

import (
	"bufio"
	"fmt"
	"os"
)

func TryHelloWorld() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter your name: ")
	name, _ := reader.ReadString('\n')
	fmt.Printf("Hello %s", name)
}
