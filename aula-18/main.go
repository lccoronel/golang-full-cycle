package main

import (
	"fmt"

	"github.com/lccoronel/golang-full-cycle/carro"
	"github.com/lccoronel/golang-full-cycle/matematica"
)

func main() {
	resultado := matematica.Soma(10, 20)

	fmt.Println("Resultado: ", resultado)

	carro := carro.Carro{Marca: "Nissan"}

	fmt.Println(carro.MostraMarca())
}
