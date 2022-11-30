package main

import "fmt"

const message = "Hello world!"

type ID int

var (
	b bool
	c int
	d string
	e float64
	f ID = 1
)

func main() {
	var meuArray [3]int
	meuArray[0] = 1
	meuArray[1] = 2
	meuArray[2] = 10

	fmt.Println(meuArray[len(meuArray)-1])

	for i, v := range meuArray {
		fmt.Printf("O Valor do indice %d é %d\n", i, v)
	}
}
