package main

import (
	"fmt"
	"time"
)

func main() {
	result := make(chan int)

	fmt.Println("Program started")

	go func() {
		fmt.Println("Anonymous goroutine started")
		time.Sleep(2 * time.Second)
		fmt.Println("Anonymous goroutine completed.")
		result <- 42
	}()

	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout fired!")
	case resultValue := <-result:
		fmt.Println("Result from anonymous goroutine:", resultValue)
	}
}
