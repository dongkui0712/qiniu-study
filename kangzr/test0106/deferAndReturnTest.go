package main

import (
	"log"
	"time"
)

func main() {
	j := DeferAndReturn()
	log.Println(time.Now().String(), "return:", j)
	//	defer fmt.Println("defer:", time.Now().String())
	//
	//	return fmt.Println("return:", time.Now().String())
}

func DeferAndReturn() int {
	i := 1
	defer func() {
		i = i + 5
		log.Println(time.Now().String(), "defer:", i)
	}()

	return i
}
