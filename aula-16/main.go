package main

import "fmt"

func main() {
	var nome interface{} = "Lucas"

	println(nome.(string))

	res, ok := nome.(int)

	fmt.Printf("valor %v, e ok e %v", res, ok)
}
