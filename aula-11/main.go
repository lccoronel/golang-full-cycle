package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	lucas := Cliente{
		Nome:  "Lucas",
		Idade: 25,
		Ativo: true,
	}

	lucas.Idade = 26

	fmt.Println(lucas)
}
