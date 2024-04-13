package main

import "fmt"

/*
Usando el ejercicio anterior, crea un programa que use “else if” y “else”.
*/
func main() {
	switch {
	case 2 == 4:
		fmt.Println("No debería imprimir")
	case 3 == 3:
		fmt.Println("Si debería imprimir")
		fallthrough
	case 4 == 5:
		fmt.Println("No debería imprimir 4 no es igual a 5")
	default:
		fmt.Println("No debería imprimir")
	}
}
