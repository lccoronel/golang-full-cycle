package main

import "fmt"

// receive only
func recebe(nome string, hello chan<- string) {
	hello <- nome
}

// read only or send only
func ler(data <-chan string) {
	fmt.Println(<-data)
}

// Thread 1
func main() {
	hello := make(chan string)
	go recebe("Hello", hello)
	ler(hello)
}
