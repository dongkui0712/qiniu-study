package main

import (
	"fmt"
)

func main() {
	fmt.Printf("this is return print %d", myReturn())
}

func myReturn() int {
	defer fmt.Println("Thist is defer print")
	return 1
}
