package main

import (
	"fmt"
	"math"
)

const Pi = 3.14
var c, python, java bool

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	i := 42
	fmt.Printf("i is of type %T\n", i)
	fmt.Println("My favorite number is", math.Sqrt(9))
	fmt.Println("0.1 + 0.2 = 0.3", 0.1 + 0.2 == 0.3)
	fmt.Println("Pi is", math.Pi)
	fmt.Println(Pi)
	fmt.Println("5 + 7 =", add(5, 7))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(c, python, java, i)
}
