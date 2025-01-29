// Go program to illustrate the
// use of simple range loop
package main

import "fmt"

// Main function
func main() {

	// Here rangoVariable is a array
	rangoVariable := []string{"GFG", "Geeks", "GeeksForGeeks"}

	// i and j stores the value of rvariable
	// i store index number of individual string and
	// j store individual string of the given array
	for i, j := range rangoVariable {
		fmt.Println(i, j)
	}

}
