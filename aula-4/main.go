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
	fmt.Printf("O tipo de E é %T", f)
}
