package main

import (
	"fmt"
	"time"
)

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}

func main() {
	data := make(chan int)
	qtdWorkers := 10

	// inicializa os workers
	for i := 0; i < qtdWorkers; i++ {
		go worker(i, data)
	}

	// request
	for i := 0; i < 100; i++ {
		data <- i
	}
}
