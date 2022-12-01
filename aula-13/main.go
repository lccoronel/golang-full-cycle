package main

import "fmt"

type Endereco struct {
	Rua    string
	Numero int
	Cidade string
	Estado string
}

type Pessoa interface {
	Desativer()
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (cliente Cliente) Desativer() {
	cliente.Ativo = false
	fmt.Printf("O cliente %s foi desativado", cliente.Nome)
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativer()
}

func main() {
	lucas := Cliente{
		Nome:  "Lucas",
		Idade: 25,
		Ativo: true,
	}

	lucas.Ativo = false
	Desativacao(lucas)
}
