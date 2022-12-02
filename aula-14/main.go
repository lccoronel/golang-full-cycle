package main

import "fmt"

type Cliente struct {
	nome string
}

func (c *Cliente) andou() {
	c.nome = "Lucas Coronel"
	fmt.Printf("O cliente %s andou\n", c.nome)
}

func main() {
	lucas := Cliente{
		nome: "Lucas",
	}

	lucas.andou()
	println(lucas.nome)
}
