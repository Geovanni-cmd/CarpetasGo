package main

import "fmt"

func main() {
	var num1, num2 int16
	fmt.Print("Ingresa el numero 1: ")
	fmt.Scan(&num1)
	fmt.Print("Ingresa el numero 1: ")
	fmt.Scan(&num2)

	var suma, resta, multi int16
	suma = num1 + num2
	fmt.Printf("El resultado de la suma entre %d + %d es: %d  \n", num1, num2, suma)
	resta = num1 - num2
	fmt.Printf("El resultado de la resta entre %d - %d es: %d  \n", num1, num2, resta)
	multi = num1 * num2
	fmt.Printf("El resultado de la multiplicacion entre %d * %d es: %d  \n", num1, num2, multi)
	div := float32(num1) / float32(num2)
	fmt.Printf("El resultado de la division entre %d / %d es: %f  \n", num1, num2, div)

}
