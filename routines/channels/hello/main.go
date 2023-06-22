package main

import "fmt"

// Thread 1
func main() {
	canal := make(chan string) // canal vazio

	// Thread 2
	go func() {
		canal <- "Ola mundo" // canal cheio
	}()

	// Thread 1
	msg := <-canal // canal vazio
	fmt.Println(msg)
}
