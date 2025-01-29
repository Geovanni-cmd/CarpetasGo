package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 60
	fmt.Println("Numero: ", num)
	fmt.Println("Decimal en binario: ", conBin(num))
}

func conBin(n int) string {
	residuo := ""
	for n > 0 {
		residuo += strconv.Itoa(n % 2)
		n = n / 2
	}

	residuoReverse := ""
	for _, c := range residuo {
		residuoReverse = string(c) + residuoReverse
	}

	return residuoReverse
}
