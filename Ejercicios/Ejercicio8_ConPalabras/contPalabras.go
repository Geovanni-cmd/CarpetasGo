/*
 * Crea un programa que cuente cuantas veces se repite cada palabra
 * y que muestre el recuento final de todas ellas.
 * - Los signos de puntuación no forman parte de la palabra.
 * - Una palabra es la misma aunque aparezca en mayúsculas y minúsculas.
 * - No se pueden utilizar funciones propias del lenguaje que
 *   lo resuelvan automáticamente.
 */

package main

import (
	"fmt"
	"strings"
)

func main() {
	cadena := "Geo geo geo"
	fmt.Println(cadena)
	fmt.Println("Palabras contadas: ", contPalabras(strings.ToLower(cadena)))

}

func contPalabras(c string) int {

	cadena := []byte(c)
	cadenaCont := ""
	contador := 0

	for i := 0; i < len(c); i++ {
		cadenaCont += string(cadena[i])
		fmt.Println("cadena 1:", cadenaCont)
		if cadena[i] == 32 || i == (len(c)-1) {
			cadenaContDos := ""
			for i := 0; i < len(c); i++ {
				cadenaContDos += string(cadena[i])
				if cadenaCont == cadenaContDos {
					fmt.Println("Cadena 2", cadenaContDos)
					fmt.Println(contador)
					contador += 1
					cadenaContDos = ""
				}

			}
		}
	}

	return contador
}
