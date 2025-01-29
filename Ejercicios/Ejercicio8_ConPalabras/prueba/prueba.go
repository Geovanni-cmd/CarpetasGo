package main

import (
	"fmt"
	"strings"
)

type cadenasCont struct {
	cadena   string
	cantidad int
}

func main() {
	cadena := "Hola mi nombre es david mi nombre david david mi"
	fmt.Println(cadena)
	cadenaMinus := strings.ToLower(cadena)
	//fmt.Println(cadenaMinus)
	cadenaBytes := []byte(cadenaMinus)
	//fmt.Println(cadenaBytes)
	nuevaCad := ""
	//cadenaF := ""

	for i := 0; i < len(cadena); i++ {
		nuevaCad += string(cadenaBytes[i])
		//fmt.Println(nuevaCad)
		if cadenaBytes[i] == 32 || i == (len(cadena)-1) {

			fmt.Printf("\nPalabra: %s esta: %d", nuevaCad, strings.Count(cadenaMinus, nuevaCad))
			nuevaCad = ""
		}
	}

}
