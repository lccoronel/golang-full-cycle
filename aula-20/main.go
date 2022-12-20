package main

func main() {
	a := 1
	b := 2
	c := 3

	if a > b {
		println(a)
	} else {
		println(b)
	}

	// consicional e
	if a > b && c > b {
		println("condicional &&")
	}

	// condicional ou
	if a > b || c > b {
		println("condicional ||")
	}

	// switch
	switch a {
	case 1:
		println("a")

	case 2:
		println("b")

	default:
		println("default")

	}
}
