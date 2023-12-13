package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("Program started...")

	done := make(chan struct{})
	go func() {
		sayHello("Hazadus")
		done <- struct{}{}
	}()
	<-done

	fmt.Println("Program completed in ", time.Since(start))
}

func sayHello(name string) {
	fmt.Println("\tHello, " + name + "!")
	time.Sleep(time.Millisecond * 500)
}
