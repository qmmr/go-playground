package main

import (
	"fmt"
	"time"
)

func printMessage(text string) {
	for i := 0; i < 10; i++ {
		fmt.Println(text)
		time.Sleep(time.Microsecond * 500)
	}
}

func main() { // this is main goroutine
	go printMessage("Go is cool!") // this creates a new goroutine
	printMessage("Go is interesting...")
}
