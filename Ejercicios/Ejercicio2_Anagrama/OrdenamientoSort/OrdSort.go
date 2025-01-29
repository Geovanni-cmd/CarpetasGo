package main

import (
	"fmt"
	"sort"
)

func sortString(s string) string {
	// Convertir el string en un slice de runas
	runes := []rune(s)

	// Ordenar el slice de runas
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	// Convertir el slice de runas ordenado de nuevo a un string
	return string(runes)
}

func main() {
	s := "botines"
	sortedString := sortString(s)
	fmt.Println("Original:", s)
	fmt.Println("Ordenado:", sortedString)
}
