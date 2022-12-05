package main

import "fmt"

func main() {
	var x interface{} = 10
	var y interface{} = "Hello"

	showType(x)
	showType(y)
}

func showType(t interface{}) {
	fmt.Printf("O tipo da vairavel eé %T e o valor eé %v\n", t, t)
}
