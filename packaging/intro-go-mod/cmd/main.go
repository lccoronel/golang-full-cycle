package main

import (
	"fmt"

	"github.com/lccoronel/intro-go-mod/math"
)

func main() {
	m := math.NewMath(1, 2)

	fmt.Println(m.Add())
}
