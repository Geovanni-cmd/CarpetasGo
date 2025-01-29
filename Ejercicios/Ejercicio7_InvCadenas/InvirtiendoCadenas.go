/*
 * Crea un programa que invierta el orden de una cadena de texto
 * sin usar funciones propias del lenguaje que lo hagan de forma automática.
 * - Si le pasamos "Hola mundo" nos retornaría "odnum aloH"
 */

package main

import "fmt"

func main() {
	fmt.Println(reversa("Hola Mundo"))
}

func reversa(s string) string {
	contador := len(s) - 1
	cadenaReversa := ""
	cadena := []byte(s)
	for i := 0; i < len(s); i++ {
		cadenaReversa += string(cadena[contador])
		contador--
	}

	return cadenaReversa
}
