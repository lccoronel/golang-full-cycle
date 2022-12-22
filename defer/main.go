package main

import "fmt"

func main() {
	fmt.Println("Linha 1")
	defer fmt.Println("Linha 2")
	fmt.Println("Linha 3")
}

// defer faz com que a linha seja executada por ultimo
