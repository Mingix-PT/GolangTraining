package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2*z)
		fmt.Println(z)
	}
	return z
}

func AdvanceSqrt(x float64) float64 {
	z := 1.0
	for {
		z -= (z*z - x) / (2*z)
		fmt.Println(z)
		if math.Abs(z*z-x) < 1e-12 {
			break
		}
	}
	return z
}

func main() {
	fmt.Println(math.Sqrt(2))
	fmt.Println(AdvanceSqrt(2))
	fmt.Println(math.Sqrt(2) - Sqrt(2))
	fmt.Println(math.Sqrt(2) - AdvanceSqrt(2))
}


