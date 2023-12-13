package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	names := [3]string{"Michael", "Dmitry", "Vladimir"}

	start := time.Now()
	fmt.Println("Program started...")

	var wg sync.WaitGroup
	for _, name := range names {
		wg.Add(1)
		//
		// Note how we explicitly pass `name` into anonymous function:
		go func(name string) {
			sayHello(name)
			wg.Done()
		}(name)
	}

	wg.Wait()
	fmt.Println("Program completed in ", time.Since(start))
}

func sayHello(name string) {
	fmt.Println("\tHello, " + name + "!")
	time.Sleep(time.Millisecond * 500)
}
