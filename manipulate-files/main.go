package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// creating
	file, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	// wrinting
	tamanho, err := file.Write([]byte("Escrevendo dados no arquivo"))
	// tamanho, err := file.WriteString("Hello, Lucas!")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes", tamanho)

	file.Close()

	// reading
	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(arquivo))

	// reading little by little
	file2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file2)
	buffer := make([]byte, 10)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	// removing
	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}
}
