package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) complex128 {
	if x < 0 {
		return complex(0, real(Sqrt(-x)))
	}
	z := 1.0
	for math.Abs(z*z - x) > 0.000001 {
		z -= (z*z - x) / (2 * z)
	}
	return complex(z, 0)
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
	fmt.Println(Sqrt(-2))
	fmt.Println(math.Sqrt(-2))
}
