package main

func soma(a, b *int) int {
	return *a + *b
}

func main() {
	a := 10
	var ponteiro *int = &a
	*ponteiro = 20
	b := &a
	println(*b)

	number1 := 2
	number2 := 3

	println(soma(&number1, &number2))
}
