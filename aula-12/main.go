package main

import "fmt"

type Endereco struct {
	Rua    string
	Numero int
	Cidade string
	Estado string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func main() {
	lucas := Cliente{
		Nome:  "Lucas",
		Idade: 25,
		Ativo: true,
	}

	lucas.Rua = "Teste"
	lucas.Numero = 23
	lucas.Cidade = "SP"
	lucas.Estado = "SP"

	fmt.Println(lucas)
}
