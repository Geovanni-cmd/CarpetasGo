package main

import "fmt"

func main() {
	m, _, s, r := calculadora(200, 3)
	fmt.Print("200 x 3 = ", m)
	fmt.Print("\n200 / 3 = ")
	fmt.Print("\n200 + 3 = ", s)
	fmt.Print("\n200 - 3 = ", r)
}

func calculadora(a, b float32) (mul, div, sum, res float32) {
	mul = a * b
	div = a / b
	sum = a + b
	res = a - b
	return
}
