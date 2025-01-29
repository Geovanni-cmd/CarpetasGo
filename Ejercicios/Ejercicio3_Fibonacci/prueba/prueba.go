package main

import "fmt"

func main() {
	n0 := 0
	n1 := 1

	for i := 1; i <= 50; i++ {
		fmt.Println(n0)

		var fib = n0 + n1
		n0 = n1
		n1 = fib
	}

}
