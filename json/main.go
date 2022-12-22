package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int `json:"n"`
	Saldo  int `json:"s"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 100}

	res, err := json.Marshal(conta)
	if err != nil {
		println(err)
	}

	println(string(res))

	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(conta)

	// para funfar tirar as tafs da struct
	jsonPuro := []byte(`{"Numero":2,"Saldo":200}`)
	var contaX Conta
	json.Unmarshal(jsonPuro, &contaX)
	println(contaX.Saldo)

	jsonDiff := []byte(`{"n":3,"s":200}`)
	var contaY Conta
	json.Unmarshal(jsonDiff, &contaY)
	println(contaY.Numero)
}

// defer faz com que a linha seja executada por ultimo
