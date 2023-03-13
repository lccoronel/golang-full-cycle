package main

import "github.com/lccoronel/go-mod-replace/math"

func main() {
	m := math.NewMath(1, 2)
	println(m.Add())
}
