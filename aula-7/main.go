package main

import "fmt"

func main() {
	salarios := map[string]int{"Lucas": 1000, "Joao": 1000, "Maria": 1000}
	fmt.Println(salarios["Lucas"])

	delete(salarios, "Maria")

	salarios["Marcos"] = 1000
	fmt.Println(salarios["Marcos"])

	// criandoMap1 := make(map[string]int)
	// criandoMap2 := map[string]int{}

	for nome, salario := range salarios {
		fmt.Printf("O salario de %s é %d\n", nome, salario)
	}
}
