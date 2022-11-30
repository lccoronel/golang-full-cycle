package main

import (
	"errors"
	"fmt"
)

func main() {
	valor, err := sum(50, 10)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(valor)
}

func sum(a int, b int) (int, error) {
	if a+b >= 50 {
		return a + b, errors.New("Soma é maior que 50")
	}

	return a + b, nil
}
