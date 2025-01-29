package main

import "fmt"

type Botella struct {
	agua      int
	coca_cola int
	lleno     bool
}

type metodos interface {
	edad() int
}

type persona struct {
	nacimiento int64
	nombre     string
	vivo       bool
}

func main() {
	botella := Botella{agua: 12, coca_cola: 2}
	fmt.Println("Agua:", botella.agua)
	fmt.Println("Coca Cola:", botella.coca_cola)
	fmt.Println("Lleno:", botella.lleno)

	var botella2 Botella
	fmt.Println(botella2)
	botella2.agua = 299
	fmt.Println(botella2)

}
