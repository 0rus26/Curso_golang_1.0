package main

import "fmt"

/*
Usando el ejercicio anterior, crea un programa que use “else if” y “else”.
*/
func main() {
	i := 102
	if i < 100 {
		fmt.Printf("%v es Menor a 100 ", i)
	} else if i == 100 {
		fmt.Printf("%v es igual a 100", i)
	} else {
		fmt.Printf("%v es Mayor a 100", i)
	}
}
