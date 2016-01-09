package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(time.Second * 2)
	quit := make(chan bool)
	go func() {
		for {
			select {
			case <-ticker.C:
				// do stuff
				fmt.Println("Timer 1 expired")

			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
	time.Sleep(25 * time.Second)
	quit <- true
	time.Sleep(25 * time.Second)

}
