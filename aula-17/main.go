package main

type MyNumber int

// "~" abrir excessoes para usar MyNumber como um tipo
type Number interface {
	~int | ~float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}

	return soma
}

// "comparable" deixar usar o generics desde que os parametros possao ser comparaveis
func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{"Lucas": 25, "Geisa": 25, "Lilian": 19}
	m2 := map[string]float64{"Lucas": 25.5, "Geisa": 25.3, "Lilian": 19.8}
	m3 := map[string]MyNumber{"Lucas": 25, "Geisa": 25, "Lilian": 19}

	println(Soma(m))
	println(Soma(m2))
	println(Soma(m3))

	println(Compara(10, 10))
}
