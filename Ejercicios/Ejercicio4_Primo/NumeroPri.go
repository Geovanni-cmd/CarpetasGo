package main

import "fmt"

func main() {

	//fmt.Println(15 % 1)
	for i := 1; i <= 100; i++ {
		if esPrimo(i) {
			fmt.Println(i)
		}
	}
}

func esPrimo(num int) bool {
	var a int
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			a++
		}
	}

	if a == 2 {
		return true
	}

	return false
}
