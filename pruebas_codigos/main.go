package main

import "fmt"

func main() {
	var name string
	var numA int
	//var numB int
	fmt.Println("Hola")
	fmt.Scanf("%s", &name)
	fmt.Printf("Nombre: %s ", name)

	fmt.Println("Ingresa el numero A: ")
	fmt.Scanf("%d", &numA)
	numA = 2
	fmt.Printf("El numero A es: %d", numA)
	/*fmt.Println("Ingresa el numero B: ")
	fmt.Scanf("%d", numB)

	fmt.Printf("La suma del numero A:%d y el numero B:%d, es: %v", numA, numB, numA+numB)*/
}
