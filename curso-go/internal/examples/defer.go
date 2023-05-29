package examples

import (
	"fmt"
	"log"
)

func TryDefer() {
	printDefers()
	tryPanic()
}

func printDefers() {
	defer fmt.Println("This should be printing second")
	fmt.Println("This should print first")
}

func tryPanic() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recovered from panic:", err)
		}
	}()
	panic(fmt.Errorf("panic"))
	fmt.Println("This should never print")
}
