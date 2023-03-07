package main

import (
	"fmt"

	"github.com/lccoronel/intro-go-mod/math"
)

func main() {
	m := math.Math{A: 2, B: 2}

	fmt.Println(m.Add())
}
