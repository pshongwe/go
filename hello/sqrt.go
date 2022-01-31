package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 3; i++ {
		z -= (z*z - x) / (2*z)
	}
	return z
}

func main() {
	fmt.Println("func:", Sqrt(2))
	fmt.Println("library", math.Sqrt(2))
}

