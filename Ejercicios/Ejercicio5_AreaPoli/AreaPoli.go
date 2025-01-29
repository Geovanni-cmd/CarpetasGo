/*
 * Crea una única función (importante que sólo sea una) que sea capaz
 * de calcular y retornar el área de un polígono.
 * - La función recibirá por parámetro sólo UN polígono a la vez.
 * - Los polígonos soportados serán Triángulo, Cuadrado y Rectángulo.
 * - Imprime el cálculo del área de un polígono de cada tipo.
 */

package main

import "fmt"

func main() {
	t := triangulo{base: 10.0, altura: 5.0}
	c := cuadrado{4}
	r := rectangulo{5, 7}

	fmt.Println("------------------------------------")
	fmt.Println("Area del cuadrado: ", c.Area())
	fmt.Println("Area del triangulo: ", t.Area())
	fmt.Println("Area del rectangulo: ", r.Area())
	fmt.Println("Area del rectangulo: ", r.Perimetro())
}

func (r rectangulo) Perimetro() float32 {
	return 2 * (r.base + r.altura)
}

func (c cuadrado) Area() float32 {
	return float32(c.base) * float32(c.base)
}
func (c rectangulo) Area() float32 {
	return float32(c.base) * float32(c.altura)
}
func (c triangulo) Area() float32 {
	return (float32(c.base) * float32(c.altura)) / 2
}

type Poligono interface {
	Area() float32
	Perimetro() float32
}

type triangulo struct {
	base   float32
	altura float32
}

type cuadrado struct {
	base float32
}

type rectangulo struct {
	base   float32
	altura float32
}
