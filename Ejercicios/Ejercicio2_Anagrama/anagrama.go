package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	esAnagrama := esAnagramaFun("botines", "bisonte")
	fmt.Print(esAnagrama)
}

func esAnagramaFun(pal1, pal2 string) bool {
	if len(pal1) != len(pal2) {
		return false
	} else {
		//Se convierten en minusculas por si tienen letras en mayusculas
		pal1 = strings.ToLower(pal1)
		pal2 = strings.ToLower(pal2)

		pal1Rune := []rune(pal1)
		pal2Rune := []rune(pal2)

		sort.Slice(pal1Rune, func(i, j int) bool {
			return pal1Rune[i] < pal1Rune[j]
		})

		sort.Slice(pal2Rune, func(i, j int) bool {
			return pal2Rune[i] < pal2Rune[j]
		})

		for i := range pal1Rune {
			fmt.Println(pal1Rune[i])
			if pal1Rune[i] != pal2Rune[i] {
				return false
			}
		}
	}
	return true

}
